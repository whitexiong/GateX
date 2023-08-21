package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

// 面板接口
func SetupDashboardRoutes(r *gin.Engine) {
	dashboard := r.Group("/dashboard")

	// 使用 Casbin 中间件进行权限验证
	dashboard.Use(middleware.InitJWTMiddleware())
	dashboard.GET("", handlers.ShowDashboard)
}
