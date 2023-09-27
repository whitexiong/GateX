package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID          string
	Conn        *websocket.Conn
	Pool        *ClientPool
	UserID      int   // 用户ID
	ChatRoomIDs []int // 该用户参与的所有聊天室
}

type ClientMessage struct {
	ChatID   string `json:"chat_id"`
	Type     int8   `json:"type"`
	Body     string `json:"body"`
	SenderID uint
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message:", err)
			return
		}

		if messageType == websocket.TextMessage {
			log.Printf("Received text message from client %s: %s", c.ID, string(p))
			var clientMsg ClientMessage
			if err := json.Unmarshal(p, &clientMsg); err != nil {
				log.Println("Error deserializing client message:", err)
				continue
			}

			c.Pool.Broadcast <- clientMsg
		} else if messageType == websocket.BinaryMessage {
			log.Println("Received a binary message of length:", len(p))
		}
	}

}
