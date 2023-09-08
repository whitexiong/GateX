package routes

import (
	"gateway/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	group := r.Group("/route")

	group.GET("list", handlers.GetAllRoutes)
	group.POST("add", handlers.CreateRoute)
	group.GET("detail/:id", handlers.GetRoute)
	group.POST("update/:id", handlers.UpdateRoute)
	group.GET("delete/:id", handlers.DeleteRoute)
	group.GET("paths/:path", handlers.GetRoutePathList)
}
