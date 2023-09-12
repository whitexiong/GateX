package routes

import "github.com/gin-gonic/gin"

var AllRoutes = []func(r *gin.Engine){
	SetupDashboardRoutes,
	SetupUserRoutes,
	SetupMenuRoutes,
	SetupRoleRoutes,
	SetupRoute,
	SetupUploads,
}
