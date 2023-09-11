package handlers

import (
	"fmt"
	"gateway/apierrors"
	"gateway/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetRole(c *gin.Context) {
	var role models.Role
	var casbinRules []models.CasbinRule

	roleID := c.Param("id")
	result := models.DB.Find(&role, roleID)
	if result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}

	models.DB.Where("v0 = ?", role.Name).Find(&casbinRules)

	var permissions []uint

	models.DB.Model(&models.CasbinRule{}).Where("v0 = ?", role.Name).Distinct("v3").Pluck("v3", &permissions)
	responseData := structs.Map(role)
	responseData["Permissions"] = permissions
	SendResponse(c, http.StatusOK, 200, responseData)
	return
}

func GetRoleList(c *gin.Context) {
	var roles []models.Role
	if err := models.DB.Find(&roles).Error; err != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}
	SendResponse(c, http.StatusOK, 200, roles)
	return
}

func AddRole(c *gin.Context) {
	var role models.RoleRequest

	if err := c.BindJSON(&role); err != nil {
		apierrors.HandleGinError(c, apierrors.InvalidRequestData)
		return
	}

	tx := models.DB.Begin()

	if err := tx.Create(&role.Role).Error; err != nil {
		tx.Rollback()
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}

	for _, routeID := range role.Permissions {
		var route models.Route
		result := tx.First(&route, routeID)
		if result.Error != nil {
			tx.Rollback()
			apierrors.HandleGinError(c, apierrors.DatabaseError)
			return
		}

		// p 策略模式 v1 路由  v2 请求方式
		obj := strings.Trim(route.Path, "/")
		casbinRule := models.CasbinRule{
			PType: "p",
			V0:    role.Role.Name,
			V1:    "/" + obj,
			V2:    "*",
			V3:    strconv.FormatUint(uint64(route.ID), 10),
		}

		if err := tx.Create(&casbinRule).Error; err != nil {
			tx.Rollback()
			apierrors.HandleGinError(c, apierrors.DatabaseError)
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)

		return
	}

	SendResponse(c, http.StatusOK, 200, role)
	return
}

func UpdateRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))
	var updatedRoleRequest models.RoleRequest

	if err := c.BindJSON(&updatedRoleRequest); err != nil {
		apierrors.HandleGinError(c, apierrors.InvalidRequestData)
		return
	}

	tx := models.DB.Begin()

	// 更新角色信息
	if err := tx.Model(&models.Role{}).Where("id = ?", roleID).Updates(updatedRoleRequest.Role).Error; err != nil {
		tx.Rollback()
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}

	// 删除旧的casbin规则
	if err := tx.Delete(&models.CasbinRule{}, "v0 = ?", updatedRoleRequest.Role.Name).Error; err != nil {
		tx.Rollback()
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}

	// 添加新的casbin规则
	for _, routeID := range updatedRoleRequest.Permissions {
		var route models.Route
		result := tx.First(&route, routeID)
		if result.Error != nil {
			tx.Rollback()
			apierrors.HandleGinError(c, apierrors.InternalServerError)
			return
		}

		// p 策略模式 v1 路由  v2 请求方式
		obj := strings.Trim(route.Path, "/")
		casbinRule := models.CasbinRule{
			PType: "p",
			V0:    updatedRoleRequest.Role.Name,
			V1:    "/" + obj,
			V2:    "*",
			V3:    strconv.FormatUint(uint64(route.ID), 10),
		}

		if err := tx.Create(&casbinRule).Error; err != nil {
			tx.Rollback()
			apierrors.HandleGinError(c, apierrors.InternalServerError)
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}

	SendResponse(c, http.StatusOK, 200, updatedRoleRequest)
	return
}

func DeleteRole(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	if err := models.DB.Delete(&models.Role{}, roleId).Error; err != nil {
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}
	SendResponse(c, http.StatusOK, 200, nil)
	return
}

func GetPermissions(c *gin.Context) {
	var routes []models.Route
	if err := models.DB.Find(&routes).Error; err != nil {
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}
	SendResponse(c, http.StatusOK, 200, routes)
	return
}

func AddPermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var routeIds []int
	if err := c.BindJSON(&routeIds); err != nil {
		apierrors.HandleGinError(c, apierrors.InternalServerError)
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
			apierrors.HandleGinError(c, apierrors.InternalServerError)
			return
		}
	}
	SendResponse(c, http.StatusOK, 200, nil)
	return
}

func GetRolePermissions(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Param("roleId"))
	var rules []models.CasbinRule
	if err := models.DB.Where("v0 = ?", strconv.Itoa(roleId)).Find(&rules).Error; err != nil {
		apierrors.HandleGinError(c, apierrors.InternalServerError)
		return
	}
	SendResponse(c, http.StatusOK, 200, rules)
}
