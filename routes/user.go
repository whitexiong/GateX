package routes

import (
	"gateway/auth"
	"gateway/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	group := r.Group("/user")
	group.POST("login", auth.Login)
	group.POST("logout", auth.Logout)
	group.GET("list", handlers.GetAllUsers)
	group.POST("add", handlers.CreateUser)
	group.GET("detail/:id", handlers.GetUserDetail)
	group.POST("update/:id", handlers.UpdateUser)
	group.GET("delete/:id", handlers.DeleteUser)
	group.GET("menus", handlers.GetUserMenus)
}
