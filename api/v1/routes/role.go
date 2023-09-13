package routes

import (
	"gateway/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(r *gin.Engine) {
	group := r.Group("/role")
	group.GET("list", handlers.GetRoleList)
	group.POST("add", handlers.AddRole)
	group.POST("update/:id", handlers.UpdateRole)
	group.GET("delete/:id", handlers.DeleteRole)
	group.GET("detail/:id", handlers.GetRole)
	group.GET("permissions", handlers.GetAllRoutes)
	group.GET("permissions/:roleId", handlers.GetRolePermissions)
}