package middleware

import "github.com/gin-gonic/gin"

func InitializeMiddlewares() []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{}

	middlewares = append(middlewares, InitLoggingAndErrorHandlingMiddleware()) //日志记录
	middlewares = append(middlewares, InitializeCustomErrorMiddleware())       // 初始化错误消息
	middlewares = append(middlewares, InitializeCorsMiddleware())              // 跨域
	middlewares = append(middlewares, InitJWTMiddleware())                     // 初始化 JWT 认证
	middlewares = append(middlewares, InitializeCasbinMiddleware())            // 初始化casbin权限

	return middlewares
}
