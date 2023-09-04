package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(r *gin.Engine) {
	group := r.Group("/role")
	group.Use(middleware.InitJWTMiddleware())
	group.GET("list", handlers.GetRoleList)
	group.POST("add", handlers.AddRole)
	group.DELETE("delete/:id", handlers.DeleteRole)
	group.PUT("update/:id", handlers.UpdatePolicy)
	group.GET("detail/:id", handlers.DetailPolicy)

	group.GET("permissions", handlers.GetAllRoutes)
	group.POST("add-permissions/:roleId", handlers.AddPermissions)
	group.GET("permissions/:roleId", handlers.GetRolePermissions)
}
