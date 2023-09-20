package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupDashboard(r *gin.Engine) {
	openvpnGroup := r.Group("/openvpn")

	openvpnGroup.GET("dashboard", openvpn.Dashboard)

	//// Log Management
	//log := openvpnGroup.Group("/log")
	//log.GET("/view", openvpn.viewLogs)
	//log.GET("/analysis", openvpn.analyzeLogs)
	//
	//// Security
	//security := openvpnGroup.Group("/security")
	//security.PUT("/twofa", openvpn.configureTwoFactorAuthentication) // Two-factor authentication setup
	//security.PUT("/access-control", openvpn.configureAccessControl)  // Access control setup
	//
	//// System
	//system := openvpnGroup.Group("/system")
	//system.PUT("/backup", openvpn.backupConfigurations)
	//system.PUT("/restore", openvpn.restoreConfigurations)
	//system.PUT("/update", openvpn.updateSystem)
	//
	//// Help & Support
	//help := openvpnGroup.Group("/help")
	//help.GET("/documentation", openvpn.viewDocumentation)
	//help.GET("/faq", openvpn.viewFAQ)
}
