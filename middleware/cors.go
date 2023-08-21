package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeCorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	// 明确允许的来源
	config.AllowOrigins = []string{"http://192.168.0.210:8080"}

	// 允许的 HTTP 方法
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	// 允许的请求头。注意我添加了"Authorization"，以处理JWT或其他认证方法。
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	// 允许客户端携带验证信息，例如cookie。如果你使用基于cookie的身份验证，这将是必需的。
	config.AllowCredentials = true

	return cors.New(config)
}
