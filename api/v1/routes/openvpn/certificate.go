package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupConfiguration(r *gin.Engine) {
	group := r.Group("/openvpn/configuration")
	group.GET("/list", openvpn.GetCertificateList)
	group.POST("/create", openvpn.CreateCertificate)
	group.GET("/delete/:id", openvpn.RevokeCertificate)
}
