package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListPolicies(c *gin.Context) {
	var policies []models.Policy
	models.DB.Find(&policies)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"policy": policies}})
}

func CreatePolicy(c *gin.Context) {
	var policy models.Policy
	if err := c.BindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": err.Error()}})
		return
	}

	if err := models.DB.Create(&policy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "data": gin.H{"error": "Failed to create policy."}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": policy})
}

func DeletePolicy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var policy models.Policy
	if err := models.DB.Delete(&policy, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Policy deleted successfully"})
}

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

func DetailPolicy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var policy models.Policy

	if err := models.DB.First(&policy, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Policy not found"})
		return
	}

	c.JSON(http.StatusOK, policy)
}
