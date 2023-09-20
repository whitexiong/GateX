package setting

import (
	"gateway/api/v1/setting"
	"github.com/gin-gonic/gin"
)

func SetupRole(r *gin.Engine) {
	group := r.Group("/setting/role")
	group.GET("list", setting.GetRoleList)
	group.POST("add", setting.AddRole)
	group.POST("update/:id", setting.UpdateRole)
	group.GET("delete/:id", setting.DeleteRole)
	group.GET("detail/:id", setting.GetRole)
	group.GET("permissions", setting.GetAllRoutes)
	group.GET("permissions/:roleId", setting.GetRolePermissions)
}
