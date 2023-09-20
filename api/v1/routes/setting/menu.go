package setting

import (
	"gateway/api/v1/setting"
	"github.com/gin-gonic/gin"
)

func SetupMenu(r *gin.Engine) {
	group := r.Group("/setting/menu")
	group.GET("list", setting.GetAllMenus)
	group.POST("add", setting.CreateMenu)
	group.GET("detail/:id", setting.GetMenu)
	group.POST("update/:id", setting.UpdateMenu)
	group.GET("delete/:id", setting.DeleteMenu)
}
