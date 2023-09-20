package openvpn

import (
	"gateway/api/v1/openvpn"
	"github.com/gin-gonic/gin"
)

func SetupUsers(r *gin.Engine) {
	group := r.Group("/openvpn/user")
	group.GET("/list", openvpn.GetUserList)
	group.POST("/create", openvpn.CreateUser)
	group.GET("/delete/:id", openvpn.DeleteUser)
}
