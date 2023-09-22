package websocket

import (
	"gateway/api/v1/setting/auth"
	"gateway/models"
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

	// 根据用户ID查询其所有的聊天室ID
	var chatRoomUser []models.ChatRoomUser
	// 假设您已经有一个全局的DB变量，代表数据库连接
	err = models.DB.Where("user_id = ?", userID).Find(&chatRoomUser).Error
	if err != nil {
		log.Printf("Error querying ChatRoomUser: %s", err)
		conn.Close()
		return
	}

	chatRoomIDs := make([]int, len(chatRoomUser))
	for i, item := range chatRoomUser {
		chatRoomIDs[i] = int(item.ChatRoomID)
	}

	clientID := uuid.New().String()
	client := &Client{
		ID:          clientID,
		Conn:        conn,
		Pool:        pool,
		UserID:      int(userID),
		ChatRoomIDs: chatRoomIDs, // 将查询得到的房间ID列表设置到客户端实例中
	}

	pool.Register <- client

	go client.Read()
}
