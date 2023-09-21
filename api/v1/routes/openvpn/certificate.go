package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupConfiguration(r *gin.Engine) {
	group := r.Group("/openvpn/certificate")
	group.GET("/list", openvpn.GetCertificateList)
	group.POST("/add", openvpn.CreateCertificate)
	group.GET("/delete/:id", openvpn.RevokeCertificate)
}
