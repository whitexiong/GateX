package handlers

import (
	"fmt"
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

	// Here, you can remove policy from role. The logic may depend on your ORM and database setup.

	c.JSON(http.StatusOK, gin.H{"message": "Policies removed successfully"})
}

func GetRoleList(c *gin.Context) {
	var roles []models.Role
	if err := models.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve roles."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func AddRole(c *gin.Context) {
	var role models.Role

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print the status after binding
	fmt.Println("Status after binding:", role.Status)

	if err := models.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role."})
		return
	}

	// Print the role after creating
	fmt.Println("Role after creating:", role)

	c.JSON(http.StatusOK, gin.H{"message": "Role added successfully", "data": role})
}

func UpdateRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))

	var updatedRole models.Role
	if err := c.BindJSON(&updatedRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(&models.Role{}).Where("id = ?", roleId).Updates(updatedRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully"})
}

func DeleteRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))

	if err := models.DB.Delete(&models.Role{}, roleId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
