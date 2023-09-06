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
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to retrieve roles.")
		return
	}
	SendResponse(c, http.StatusOK, 200, roles, "Success")
	return
}

func AddRole(c *gin.Context) {
	var role models.RoleRequest

	if err := c.BindJSON(&role); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, err.Error())
		return
	}

	tx := models.DB.Begin() // 开始事务

	// 创建角色
	if err := tx.Create(&role.Role).Error; err != nil {
		tx.Rollback()
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to create role.")
		return
	}

	for _, routeID := range role.Permissions {
		var route models.Route
		result := tx.First(&route, routeID)
		if result.Error != nil {
			tx.Rollback()
			SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch route.")
			return
		}

		obj := strings.Trim(route.Path, "/")
		casbinRule := models.CasbinRule{
			PType: "p",
			V0:    role.Role.Name,
			V1:    "/" + obj,
			V2:    "*",
		}

		if err := tx.Create(&casbinRule).Error; err != nil {
			tx.Rollback()
			SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to create casbin rule.")
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to commit transaction.")
		return
	}

	SendResponse(c, http.StatusOK, 200, role, "Role added successfully")
	return
}

func UpdateRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var updatedRole models.Role
	if err := c.BindJSON(&updatedRole); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, err.Error())
		return
	}

	if err := models.DB.Model(&models.Role{}).Where("id = ?", roleId).Updates(updatedRole).Error; err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to update role.")
		return
	}

	SendResponse(c, http.StatusOK, 200, nil, "Role updated successfully")
	return
}

func DeleteRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	if err := models.DB.Delete(&models.Role{}, roleId).Error; err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to delete role.")
		return
	}
	SendResponse(c, http.StatusOK, 200, nil, "Role deleted successfully")
	return
}

func GetPermissions(c *gin.Context) {
	var routes []models.Route
	if err := models.DB.Find(&routes).Error; err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to retrieve permissions.")
		return
	}
	SendResponse(c, http.StatusOK, 200, routes, "Success")
	return
}

func AddPermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var routeIds []int
	if err := c.BindJSON(&routeIds); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, err.Error())
		return
	}
	for _, routeId := range routeIds {
		permission := fmt.Sprintf("role_%d_route_%d", roleId, routeId)
		rule := models.CasbinRule{
			PType: "p",
			V0:    strconv.Itoa(roleId),
			V1:    permission,
		}
		if err := models.DB.Create(&rule).Error; err != nil {
			SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to add permission.")
			return
		}
	}
	SendResponse(c, http.StatusOK, 200, nil, "Permissions added successfully")
	return
}

func GetRolePermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var rules []models.CasbinRule
	if err := models.DB.Where("v0 = ?", strconv.Itoa(roleId)).Find(&rules).Error; err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to retrieve role permissions.")
		return
	}
	SendResponse(c, http.StatusOK, 200, rules, "Success")
}
