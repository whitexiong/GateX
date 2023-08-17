package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

// 面板接口
func SetupDashboardRoutes(r *gin.Engine) {
	dashboard := r.Group("/dashboard")
	dashboard.Use(middleware.Authorize("/dashboard", "read"))
	{
		dashboard.GET("/", handlers.ShowDashboard)
	}
}
