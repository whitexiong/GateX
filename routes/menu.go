package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupMenuRoutes(r *gin.Engine) {
	menuGroup := r.Group("/menu")

	menuGroup.Use(middleware.InitJWTMiddleware())

	menuGroup.GET("list", handlers.GetAllMenus)
	//menuGroup.POST("/", handlers.CreateMenu)
	//menuGroup.GET("/:id", handlers.GetMenu)
	//menuGroup.PUT("/:id", handlers.UpdateMenu)
	//menuGroup.DELETE("/:id", handler/s.DeleteMenu)

}
