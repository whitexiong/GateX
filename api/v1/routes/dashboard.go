package routes

import (
	"gateway/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

func SetupDashboardRoutes(r *gin.Engine) {
	dashboard := r.Group("/dashboard")
	dashboard.GET("", handlers.ShowDashboard)
}
