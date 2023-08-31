package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPolicies(c *gin.Context) {
	var rules []models.CasbinRule

	if err := models.DB.Find(&rules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"policy": rules})
}

func AddPolicy(c *gin.Context) {
	var rule models.CasbinRule

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&rule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rule})
}

// ... 可以添加其他处理函数，例如删除策略、更新策略等 ...
