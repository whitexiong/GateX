package routes

import (
	"gateway/api/v1/routes/chat"
	"gateway/api/v1/routes/openvpn"
	"gateway/api/v1/routes/setting"
	"github.com/gin-gonic/gin"
)

var AllRoutes = []func(r *gin.Engine){
	setting.SetupUser,
	setting.SetupMenu,
	setting.SetupRole,
	setting.SetupRoute,

	chat.SetUpChat,

	openvpn.SetupDashboard,
	openvpn.SetupCertificates,
	openvpn.SetupConfiguration,
	openvpn.SetupUsers,
	openvpn.SetupNetwork,

	SetupDashboardRoutes,
	SetupUploads,
}
