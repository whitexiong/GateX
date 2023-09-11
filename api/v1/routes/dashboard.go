package routes

import (
	"gateway/api/v1/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupDashboardRoutes(r *gin.Engine) {
	dashboard := r.Group("/dashboard")

	// 使用 Casbin 中间件进行权限验证
	dashboard.Use(middleware.InitJWTMiddleware())
	dashboard.GET("", handlers.ShowDashboard)
}
