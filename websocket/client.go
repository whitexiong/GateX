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
	ChatRoomIDs []int // 新增：客户端所在的聊天室ID列表
}

type ClientMessage struct {
	ChatID   string `json:"chat_id"`
	Type     string `json:"type"`
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
			var clientMsg struct {
				ChatID string `json:"chat_id"`
				Type   string `json:"type"`
				Body   string `json:"body"`
			}
			if err := json.Unmarshal(p, &clientMsg); err != nil {
				log.Println("Error deserializing client message:", err)
				continue
			}

			message := ClientMessage{ChatID: clientMsg.ChatID, Type: clientMsg.Type, Body: clientMsg.Body}
			c.Pool.Broadcast <- message

		} else if messageType == websocket.BinaryMessage {
			log.Println("Received a binary message of length:", len(p))
		}
	}
}
