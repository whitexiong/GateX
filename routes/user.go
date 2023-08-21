package routes

import (
	"gateway/auth"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")

	// 登录的接口是公开权限
	//注册中间件
	//userGroup.Use(
	//middleware.JWTMiddleware(),
	//middleware.InitializeCasbinMiddleware(),
	//)

	// 登录
	userGroup.POST("/login", auth.Login)

	userGroup.POST("/logout", auth.Logout)

	// 注册
	//userGroup.POST("/register", )

	// 获取用户信息
	//userGroup.GET("/info", middleware.JWTAuth(), handlers.GetUserInfo)

	// 其他用户相关的路由可以继续在这里添加...
}
