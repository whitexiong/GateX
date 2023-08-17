package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	// 获取数据库中的所有API统计数据
	var apis []models.APIEndpoint
	models.DB.Find(&apis)

	// 返回 JSON 数据
	c.JSON(200, gin.H{
		"apis": apis,
	})
}
