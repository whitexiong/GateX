package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupMenuRoutes(r *gin.Engine) {
	group := r.Group("/menu")

	group.Use(middleware.InitJWTMiddleware())

	group.GET("list", handlers.GetAllMenus)
	group.POST("add", handlers.CreateMenu)
	group.GET("detail/:id", handlers.GetMenu)
	group.POST("update/:id", handlers.UpdateMenu)
	group.GET("delete/:id", handlers.DeleteMenu)
}
