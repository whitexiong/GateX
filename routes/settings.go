package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupSettingsRoutes(r *gin.Engine) {
	settings := r.Group("/settings")

	// 仅允许有“read”权限的用户查看设置
	settings.GET("/", middleware.Authorize("/settings/view", "read"), handlers.ShowDashboard)

	// 仅允许有“write”权限的用户修改设置
	//settings.PUT("/", middleware.Authorize("/settings/edit", "write"), handlers.EditSettings)

	// 仅允许有特殊的“reset”权限的用户重置设置
	//settings.POST("/reset", middleware.Authorize("/settings/reset", "reset"), handlers.ResetSettings)
}
