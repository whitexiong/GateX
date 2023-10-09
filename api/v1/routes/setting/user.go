package setting

import (
	"gateway/api/v1/setting"
	"gateway/api/v1/setting/auth"
	"github.com/gin-gonic/gin"
)

func SetupUser(r *gin.Engine) {
	group := r.Group("/setting/user")
	group.POST("login", auth.Login)
	group.POST("logout", auth.Logout)
	group.POST("list", setting.GetAllUsers)
	group.POST("add", setting.CreateUser)
	group.GET("detail/:id", setting.GetUserDetail)
	group.POST("update/:id", setting.UpdateUser)
	group.GET("delete/:id", setting.DeleteUser)
	group.GET("menus", setting.GetUserMenus)
}
