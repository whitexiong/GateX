package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 为角色分配策略
func AssignPolicyToRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var role models.Role
	models.DB.Preload("Policies").First(&role, roleId)

	var policies []models.Policy
	if err := c.BindJSON(&policies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, policy := range policies {
		role.Policies = append(role.Policies, policy)
	}

	models.DB.Save(&role)
	c.JSON(http.StatusOK, gin.H{"message": "Policies assigned successfully"})
}

// 移除角色的策略
func RemovePolicyFromRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var role models.Role
	models.DB.Preload("Policies").First(&role, roleId)

	var policies []models.Policy
	if err := c.BindJSON(&policies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here, you can remove policies from role. The logic may depend on your ORM and database setup.

	c.JSON(http.StatusOK, gin.H{"message": "Policies removed successfully"})
}
