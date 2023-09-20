package openvpn

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	// 获取数据库中的所有API统计数据
	var apis []models.APIEndpoint
	models.DB.Find(&apis)

	c.JSON(200, gin.H{
		"apis": apis,
	})
}
