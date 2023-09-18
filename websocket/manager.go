package websocket

import (
	"encoding/json"
	"gateway/models"
	"log"
)

type ClientPool struct {
	Clients    map[int]*Client // 修改这里
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan ClientMessage
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[int]*Client), // 以及这里
		Broadcast:  make(chan ClientMessage),
	}
}

func (pool *ClientPool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client.UserID] = client
			log.Printf("新的客户端已连接: UserID = %d\n", client.UserID)
			log.Printf("连接池大小: %d\n", len(pool.Clients))
			log.Printf("当前所有连接的客户端: %+v\n", pool.Clients)
		case client := <-pool.Unregister:
			delete(pool.Clients, client.UserID)
			log.Println("客户端已断开连接")
			log.Printf("连接池大小: %d", len(pool.Clients))
		case message := <-pool.Broadcast:
			log.Printf("广播消息: %s", message.Body)

			var chatMsg models.Message
			if err := json.Unmarshal([]byte(message.Body), &chatMsg); err != nil {
				log.Println("消息反序列化错误:", err)
				continue
			}

			if chatMsg.ChatRoomID == 0 {
				chatRoom := models.ChatRoom{
					Name:     "私人聊天室",
					RoomType: 1,
				}
				if err := models.DB.Create(&chatRoom).Error; err != nil {
					log.Println("Error creating chat room:", err)
					continue
				}
				chatMsg.ChatRoomID = chatRoom.ID

				models.DB.Create(&models.ChatRoomUser{ChatRoomID: chatRoom.ID, UserID: chatMsg.SenderID})

				// 检查 ToUserID 是否为 nil
				if chatMsg.ToUserID != nil {
					models.DB.Create(&models.ChatRoomUser{ChatRoomID: chatRoom.ID, UserID: *chatMsg.ToUserID})
				} else {
					log.Println("Warning: chatMsg.ToUserID is nil.")
					continue
				}
			}

			dbMessage := models.Message{
				SenderID:   chatMsg.SenderID,
				ChatRoomID: chatMsg.ChatRoomID,
				ToUserID:   chatMsg.ToUserID,
				Content:    chatMsg.Content,
			}
			if err := models.DB.Create(&dbMessage).Error; err != nil {
				log.Println("Error saving message to database:", err)
				continue
			}

			for userID, client := range pool.Clients {
				log.Printf("客户端: %d, 发送到：%d\n", userID, chatMsg.ToUserID)
				if chatMsg.ToUserID != nil && userID == int(*chatMsg.ToUserID) {
					if err := client.Conn.WriteJSON(chatMsg); err != nil {
						log.Println("错误:", err)
						return
					}
				}
			}
		}
	}
}
