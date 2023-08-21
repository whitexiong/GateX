package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取所有策略
func ListPolicies(c *gin.Context) {
	var policies []models.Policy
	models.DB.Find(&policies)
	c.JSON(http.StatusOK, gin.H{"policies": policies})
}

// 创建新策略
func CreatePolicy(c *gin.Context) {
	var policy models.Policy
	if err := c.BindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&policy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, policy)
}

// 删除策略
func DeletePolicy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var policy models.Policy
	if err := models.DB.Delete(&policy, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Policy deleted successfully"})
}

// 修改策略
func UpdatePolicy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var policy models.Policy

	if err := models.DB.First(&policy, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Policy not found"})
		return
	}

	if err := c.BindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&policy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, policy)
}
