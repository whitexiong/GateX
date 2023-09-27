package websocket

import (
	"encoding/json"
	"gateway/models"
	"gateway/websocket/ai"
	"log"
)

type ClientPool struct {
	Clients    map[int][]*Client // 将每个房间的ID映射到一个客户端列表
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan ClientMessage
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[int][]*Client),
		Broadcast:  make(chan ClientMessage),
	}
}

func (pool *ClientPool) Start() {
	for {
		select {
		case client := <-pool.Register:
			for _, roomID := range client.ChatRoomIDs {
				if _, ok := pool.Clients[roomID]; !ok {
					pool.Clients[roomID] = []*Client{}
				}
				pool.Clients[roomID] = append(pool.Clients[roomID], client)
			}
			log.Printf("新用户连接: UserID = %d", client.UserID)
			log.Printf("连接池客户数量: %d", len(pool.Clients))
			log.Printf("当前连接的所有客户端: %+v", pool.Clients)

		case client := <-pool.Unregister:
			for _, roomID := range client.ChatRoomIDs {
				clientsInRoom := pool.Clients[roomID]
				for i, c := range clientsInRoom {
					if c.UserID == client.UserID {
						pool.Clients[roomID] = append(clientsInRoom[:i], clientsInRoom[i+1:]...)
						break
					}
				}
				if len(pool.Clients[roomID]) == 0 {
					delete(pool.Clients, roomID)
				}
			}
			log.Println("用户断开连接")
			log.Printf("连接池客户数量: %d", len(pool.Clients))

		case message := <-pool.Broadcast:
			log.Printf("接收到的消息: %s", message.Body)

			var chatMsg models.Message
			if err := json.Unmarshal([]byte(message.Body), &chatMsg); err != nil {
				log.Printf("消息反序列化失败: %s", message.Body)
				log.Printf("错误信息: %s", err)
				continue
			}

			if chatMsg.ChatRoomID == 0 {
				chatRoom := models.ChatRoom{
					Name:     "私人聊天室",
					RoomType: 1,
				}
				if err := models.DB.Create(&chatRoom).Error; err != nil {
					log.Printf("创建聊天室失败: %s", err)
					continue
				}
				chatMsg.ChatRoomID = chatRoom.ID

				models.DB.Create(&models.ChatRoomUser{ChatRoomID: chatRoom.ID, UserID: chatMsg.SenderID})
			}

			dbMessage := models.Message{
				SenderID:   chatMsg.SenderID,
				ChatRoomID: chatMsg.ChatRoomID,
				Content:    chatMsg.Content,
			}
			if err := models.DB.Create(&dbMessage).Error; err != nil {
				log.Printf("数据库保存消息失败: %s", err)
				continue
			}

			clientsInRoom, ok := pool.Clients[int(chatMsg.ChatRoomID)]
			if !ok {
				log.Println("目标房间内无客户")
				continue
			}

			for _, client := range clientsInRoom {
				if err := client.Conn.WriteJSON(chatMsg); err != nil {
					log.Printf("发送消息失败: %s", err)
					continue
				}
			}

			log.Printf("chatMsg: %+v", chatMsg)

			var chatRoom models.ChatRoom
			if err := models.DB.Where("id = ?", chatMsg.ChatRoomID).First(&chatRoom).Error; err != nil {
				log.Printf("查询聊天室失败: %s", err)
				continue
			}

			// 如果是AI聊天室
			if chatRoom.RoomType == models.AIChatRoom {
				// 查询所有该聊天室中的用户
				var chatRoomUsers []models.ChatRoomUser
				if err := models.DB.Where("chat_room_id = ?", chatMsg.ChatRoomID).Find(&chatRoomUsers).Error; err != nil {
					log.Printf("查询聊天室用户失败: %s", err)
					continue
				}

				// 找到除AI之外的其他用户
				var realUserID uint
				for _, user := range chatRoomUsers {
					if user.UserID != chatMsg.SenderID { // 假设消息发送者是AI，那么我们需要的是另一个用户的ID
						realUserID = user.UserID
						break
					}
				}

				if realUserID == 0 {
					log.Println("无法找到真实用户的ID")
					continue
				}

				aiResponse := ai.StartChatWithAI(chatMsg.ChatRoomID, chatMsg.Content)
				aiMessage := models.Message{
					SenderID:   realUserID, // 使用真实用户的ID作为发送者
					ChatRoomID: chatMsg.ChatRoomID,
					Content:    aiResponse,
					Type:       models.AIToOne,
					AIProvider: models.XunFei,
				}

				if err := models.DB.Create(&aiMessage).Error; err != nil {
					log.Printf("数据库保存AI消息失败: %s", err)
					continue
				}

				for _, client := range clientsInRoom {
					if err := client.Conn.WriteJSON(aiMessage); err != nil {
						log.Printf("发送AI消息失败: %s", err)
						continue
					}
				}
			}

		}
	}
}
