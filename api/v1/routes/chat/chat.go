package chat

import (
	"gateway/api/v1/chat"
	"github.com/gin-gonic/gin"
)

func SetUpChat(r *gin.Engine) {
	group := r.Group("/chat")
	group.GET("/users", chat.GetChatUserList)
	group.GET("/history/:otherUserId", chat.GetChatHistoryForUser)
	group.GET("/rooms/:roomId", chat.GetChatWindow)
	group.POST("/create", chat.CreateChatRoom)
	group.GET("/delete/:roomId", chat.DeleteChatWindow)
}
