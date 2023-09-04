package handlers

import (
	"fmt"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

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
	var role models.RoleRequest

	if err := c.BindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := models.DB.Begin() // 开始事务

	// 创建角色
	if err := tx.Create(&role.Role).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role."})
		return
	}

	for _, routeID := range role.Permissions {
		var route models.Route
		result := tx.First(&route, routeID)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route."})
			return
		}

		// 直接使用完整路由作为对象
		obj := strings.Trim(route.Path, "/")

		casbinRule := models.CasbinRule{
			PType: "p",
			V0:    role.Role.Name,
			V1:    "/" + obj,
			V2:    "*", // 默认 * POST GET UPDATE DELETE 根据实际情况修改
		}

		if err := tx.Create(&casbinRule).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create casbin rule."})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction."})
		return
	}

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

func GetPermissions(c *gin.Context) {
	var routes []models.Route
	if err := models.DB.Find(&routes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve permissions."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"permissions": routes,
	})
}

func AddPermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))

	// Assume that the front-end sends the permissions as an array of route IDs.
	var routeIds []int
	if err := c.BindJSON(&routeIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For simplicity, we'll use the roleId and routeId to form a permission string and add to the casbin_rule table.
	for _, routeId := range routeIds {
		permission := fmt.Sprintf("role_%d_route_%d", roleId, routeId)
		rule := models.CasbinRule{
			PType: "p",
			V0:    strconv.Itoa(roleId),
			V1:    permission,
		}
		if err := models.DB.Create(&rule).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add permission."})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permissions added successfully"})
}

func GetRolePermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))

	var rules []models.CasbinRule
	if err := models.DB.Where("v0 = ?", strconv.Itoa(roleId)).Find(&rules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve role permissions."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"permissions": rules,
	})
}
