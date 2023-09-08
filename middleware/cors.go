package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeCorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true

	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	config.AllowCredentials = true

	return cors.New(config)
}
