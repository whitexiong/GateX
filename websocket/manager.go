package websocket

import (
	"encoding/json"
	"gateway/models"
	"log"
)

type ClientPool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan ClientMessage
}

func NewClientPool() *ClientPool {
	return &ClientPool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan ClientMessage),
	}
}

func (pool *ClientPool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Println("新的客户端已连接")
			log.Printf("连接池大小: %d", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
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
					Name:      "私人聊天室",
					IsPrivate: true,
				}
				if err := models.DB.Create(&chatRoom).Error; err != nil {
					log.Println("Error creating chat room:", err)
					continue
				}
				chatMsg.ChatRoomID = chatRoom.ID

				models.DB.Create(&models.ChatRoomUser{ChatRoomID: chatRoom.ID, UserID: chatMsg.SenderID})
				models.DB.Create(&models.ChatRoomUser{ChatRoomID: chatRoom.ID, UserID: chatMsg.ToUserID})
			}

			dbMessage := models.Message{
				SenderID:   chatMsg.SenderID,
				ChatRoomID: chatMsg.ChatRoomID,
				ToUserID:   chatMsg.ToUserID, // Added this line
				Content:    chatMsg.Content,
			}
			if err := models.DB.Create(&dbMessage).Error; err != nil {
				log.Println("Error saving message to database:", err)
				continue
			}

			//for client := range pool.Clients {
			//	// 只发送给在目标聊天室内的客户端
			//	if contains(client.ChatRoomIDs, int(chatMsg.ChatRoomID)) {
			//		if err := client.Conn.WriteJSON(chatMsg); err != nil {
			//			log.Println("错误:", err)
			//			return
			//		}
			//	}
			//}

			for client := range pool.Clients {
				err := client.Conn.WriteJSON(chatMsg)
				if err != nil {
					return
				}
				//if err := client.Conn.WriteJSON(chatMsg); err {
				//	log.Println("错误:", err)
				//	return
				//}
			}

		}
	}
}

func contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
