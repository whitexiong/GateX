package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupNetwork(r *gin.Engine) {
	group := r.Group("/openvpn/network")
	group.GET("/settings", openvpn.GetNetworkSettings)
	group.POST("/update/:id", openvpn.UpdateNetworkSettings)
}
