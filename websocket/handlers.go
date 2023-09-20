package websocket

import (
	"gateway/api/v1/setting/auth"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var secretKey = os.Getenv("JWT_SECRET_KEY")

func HandleWebSocketConnection(pool *ClientPool, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		log.Println("Token not provided in WebSocket request")
		conn.Close()
		return
	}

	userID, err := auth.JwtService.ParseToken(tokenString)
	if err != nil {
		log.Printf("Error parsing JWT: %s", err)
		conn.Close()
		return
	}

	clientID := uuid.New().String()
	client := &Client{
		ID:     clientID,
		Conn:   conn,
		Pool:   pool,
		UserID: int(userID), // Set UserID from JWT
	}

	// Register the new Client to the pool
	pool.Register <- client

	// Start reading messages for this client
	go client.Read()
}
