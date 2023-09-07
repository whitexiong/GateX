package routes

import (
	"gateway/handlers"
	"github.com/gin-gonic/gin"
)

func SetupMenuRoutes(r *gin.Engine) {
	group := r.Group("/menu")
	group.GET("list", handlers.GetAllMenus)
	group.POST("add", handlers.CreateMenu)
	group.GET("detail/:id", handlers.GetMenu)
	group.POST("update/:id", handlers.UpdateMenu)
	group.GET("delete/:id", handlers.DeleteMenu)
}
