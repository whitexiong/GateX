package routes

import (
	"gateway/api/v1/chat"
	"github.com/gin-gonic/gin"
)

func SetUpChatRoutes(r *gin.Engine) {
	group := r.Group("/chat")
	group.GET("/users", chat.GetChatUserList)
	group.GET("/history/:otherUserId", chat.GetChatHistoryForUser)
	//group.GET("/rooms/:roomId", chat.GetChatWindow)       // 获取聊天窗口的接口
	//group.POST("/rooms", chat.CreateChatRoom)             // 创建聊天房间的接口
	//group.DELETE("/rooms/:roomId", chat.DeleteChatWindow) // 删除聊天窗口的接口
}
