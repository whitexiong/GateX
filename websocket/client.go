// websocket/client.go

package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *ClientPool
}

type ClientMessage struct {
	Type string `json:"type"`
	Body string `json:"body"`
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
			message := ClientMessage{Type: "text", Body: string(p)}
			c.Pool.Broadcast <- message
		} else if messageType == websocket.BinaryMessage {
			// 这里你可以处理二进制消息。例如，你可能想保存一个接收到的图像或文件。
			// 为了简化，我们仅将其记录为二进制数据
			log.Println("Received a binary message of length:", len(p))
		}
	}
}
