package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocketConnection(pool *ClientPool, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 创建一个新的Client
	client := &Client{
		ID:   "SomeUniqueID", // 实际上，你应该为每个客户端生成一个唯一的ID
		Conn: conn,
		Pool: pool,
	}

	// 将新的Client注册到pool
	pool.Register <- client

	// 开始为该客户端监听消息
	go client.Read()
}
