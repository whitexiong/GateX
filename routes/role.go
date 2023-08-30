package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(r *gin.Engine) {
	menuGroup := r.Group("/role")
	menuGroup.Use(middleware.InitJWTMiddleware())
	menuGroup.GET("list", handlers.GetRoleList)
	menuGroup.POST("add", handlers.AddRole)
	menuGroup.DELETE("delete/:id", handlers.DeleteRole)
	menuGroup.PUT("update/:id", handlers.UpdatePolicy)
	menuGroup.GET("detail/:id", handlers.DetailPolicy)
}
