package routes

import (
	"gateway/api/v1/chat"
	"github.com/gin-gonic/gin"
)

func SetUpChatRoutes(r *gin.Engine) {
	group := r.Group("/chat")
	group.GET("/users", chat.GetChatUserList)
	group.GET("/history/:otherUserId", chat.GetChatHistoryForUser)
}
