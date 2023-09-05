package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	group := r.Group("/route")

	group.Use(middleware.InitJWTMiddleware())

	group.GET("list", handlers.GetAllRoutes)
	group.POST("add", handlers.CreateRoute)
	group.GET("detail/:id", handlers.GetRoute)
	group.POST("update/:id", handlers.UpdateRoute)
	group.GET("delete/:id", handlers.DeleteRoute)
	//menuGroup.POST("/", handlers.CreateMenu)
	//menuGroup.GET("/:id", handlers.GetMenu)
	//menuGroup.PUT("/:id", handlers.UpdateMenu)
	//menuGroup.DELETE("/:id", handler/s.DeleteMenu)

}
