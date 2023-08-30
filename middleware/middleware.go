package middleware

import "github.com/gin-gonic/gin"

func InitializeMiddlewares() []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{}

	// 添加 Casbin 中间件
	middlewares = append(middlewares, InitializeCustomErrorMiddleware())
	middlewares = append(middlewares, InitializeCorsMiddleware())
	middlewares = append(middlewares, InitializeCasbinMiddleware())
	//middlewares = append(middlewares, JWTMiddleware())

	return middlewares
}
