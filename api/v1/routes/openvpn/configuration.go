package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupCertificates(r *gin.Engine) {
	group := r.Group("/openvpn/certificate")
	group.GET("/list", openvpn.GetConfigurationList)
	group.POST("/update/:id", openvpn.UpdateConfiguration)
}
