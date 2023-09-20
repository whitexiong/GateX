package setting

import (
	"gateway/api/v1/setting"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	group := r.Group("/setting/route")

	group.GET("list", setting.GetAllRoutes)
	group.POST("add", setting.CreateRoute)
	group.GET("detail/:id", setting.GetRoute)
	group.POST("update/:id", setting.UpdateRoute)
	group.GET("delete/:id", setting.DeleteRoute)
	group.GET("paths/:path", setting.GetRoutePathList)
}
