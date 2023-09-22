package websocket

import (
	"encoding/json"
	"gateway/models"
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
			log.Printf("新的客户端已连接: UserID = %d\n", client.UserID)
			log.Printf("连接池大小: %d\n", len(pool.Clients))
			log.Printf("当前所有连接的客户端: %+v\n", pool.Clients)

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

			}

			dbMessage := models.Message{
				SenderID:   chatMsg.SenderID,
				ChatRoomID: chatMsg.ChatRoomID,
				Content:    chatMsg.Content,
			}
			if err := models.DB.Create(&dbMessage).Error; err != nil {
				log.Println("Error saving message to database:", err)
				continue
			}

			// 获取目标房间的所有客户端
			clientsInRoom, ok := pool.Clients[int(chatMsg.ChatRoomID)]
			if !ok {
				log.Println("没有客户端在目标房间")
				return
			}

			for _, client := range clientsInRoom {
				// 发送消息给目标房间的所有客户端
				log.Println("向目标房间的客户端发送消息:", client.ID)
				if err := client.Conn.WriteJSON(chatMsg); err != nil {
					log.Println("错误:", err)
					return
				}
			}

		}
	}
}
