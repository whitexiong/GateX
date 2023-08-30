package routes

import (
	"gateway/handlers"
	"github.com/gin-gonic/gin"
)

func SetupSettingsRoutes(r *gin.Engine) {
	settings := r.Group("/settings")

	// 使用 Casbin 中间件进行权限验证
	//settings.Use(middleware.InitializeCasbinMiddleware())

	settings.GET("/", handlers.ShowDashboard)
	//settings.PUT("/", handlers.EditSettings)
	//settings.POST("/reset", handlers.ResetSettings)
}
