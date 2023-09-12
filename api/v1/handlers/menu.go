package handlers

import (
	"fmt"
	"gateway/apierrors"
	"gateway/models"
	"gateway/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllMenus(c *gin.Context) {
	var menus []models.Menu
	models.DB.Order("`Order`").Find(&menus)
	transformedMenus := util.ConvertToTree(menus, util.MapMenuToTreeItem)
	SendResponse(c, http.StatusOK, 200, transformedMenus)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		SendResponse(c, http.StatusOK, apierrors.AuthenticationFailed, nil)
		apierrors.HandleGinError(c, apierrors.DataNotFound)
		return
	}
	if result := models.DB.Create(&menu); result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DataNotFound)
		return
	}
	SendResponse(c, http.StatusOK, 200, menu)
}

func GetMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DataNotFound)
		return
	}
	SendResponse(c, http.StatusOK, 200, menu)
}

func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		SendResponse(c, http.StatusOK, apierrors.InvalidRequestData, nil)
		return
	}
	if result := models.DB.Save(&menu); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}
	SendResponse(c, http.StatusOK, apierrors.Success, menu)
}

func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}
	if result := models.DB.Delete(&menu); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}
	SendResponse(c, http.StatusOK, apierrors.Success, nil)
}

func GetUserMenus(c *gin.Context) {
	// 获取用户名从 JWT 中
	roleName, exists := c.Get("role")
	if !exists {
		SendResponse(c, http.StatusOK, apierrors.Unauthorized, nil)
		return
	}

	if roleName == "root" {
		var allMenus []models.Menu
		if err := models.DB.Find(&allMenus).Error; err != nil {
			SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
			return
		}
		transformedMenus := util.ConvertToTree(allMenus, util.MapMenuToTreeItem)
		SendResponse(c, http.StatusOK, apierrors.Success, transformedMenus)
		return
	}

	role, _ := c.Get("role")
	if role == nil {
		SendResponse(c, http.StatusOK, apierrors.Unauthorized, nil)
		return
	}

	menus, err := fetchUserMenus(fmt.Sprintf("%v", role))
	if err != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	transformedMenus := util.ConvertToTree(menus, util.MapMenuToTreeItem)
	SendResponse(c, http.StatusOK, apierrors.Success, transformedMenus)
}

func fetchUserMenus(roleName string) ([]models.Menu, error) {
	var casbinRules []models.CasbinRule
	if err := models.DB.Where("v0 = ?", roleName).Find(&casbinRules).Error; err != nil {
		return nil, err
	}

	var routeIds []uint
	for _, cr := range casbinRules {
		routeId, _ := strconv.ParseUint(cr.V3, 10, 32)
		routeIds = append(routeIds, uint(routeId))
	}

	var menus []models.Menu
	if err := models.DB.Where("route_id IN ?", routeIds).Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
