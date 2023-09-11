package handlers

import (
	"fmt"
	"gateway/apierrors"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllMenus(c *gin.Context) {
	var menus []models.Menu
	models.DB.Order("`Order`").Find(&menus)

	transformedMenus := ConvertMenusToTree(menus)
	SendResponse(c, http.StatusOK, 200, transformedMenus)
}

func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
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
		apierrors.HandleGinError(c, apierrors.DataNotFound)
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		apierrors.HandleGinError(c, apierrors.InvalidRequestData)
		return
	}
	if result := models.DB.Save(&menu); result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}
	SendResponse(c, http.StatusOK, apierrors.Success, menu)
}

func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}
	if result := models.DB.Delete(&menu); result.Error != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}
	SendResponse(c, http.StatusOK, apierrors.Success, nil)
}

func GetUserMenus(c *gin.Context) {
	// 获取用户名从 JWT 中
	roleName, exists := c.Get("role")
	if !exists {
		apierrors.HandleGinError(c, apierrors.Unauthorized)
		return
	}

	if roleName == "root" {
		var allMenus []models.Menu
		if err := models.DB.Find(&allMenus).Error; err != nil {
			apierrors.HandleGinError(c, apierrors.DatabaseError)
			return
		}
		transformedMenus := ConvertMenusToTree(allMenus)
		SendResponse(c, http.StatusOK, apierrors.Success, transformedMenus)
		return
	}

	role, _ := c.Get("role")
	if role == nil {
		apierrors.HandleGinError(c, apierrors.Unauthorized)
		return
	}

	menus, err := fetchUserMenus(fmt.Sprintf("%v", role))
	if err != nil {
		apierrors.HandleGinError(c, apierrors.DatabaseError)
		return
	}

	transformedMenus := ConvertMenusToTree(menus)
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

func ConvertMenusToTree(menus []models.Menu) []map[string]interface{} {
	var transformedMenus []map[string]interface{}
	menuMap := make(map[uint]*map[string]interface{})

	for _, menu := range menus {
		transformedMenu := map[string]interface{}{
			"id":       menu.ID,
			"name":     menu.Name,
			"value":    menu.ID,
			"label":    menu.Name,
			"path":     menu.Path,
			"icon":     menu.Icon,
			"status":   menu.Status,
			"children": []map[string]interface{}{},
		}
		menuMap[menu.ID] = &transformedMenu
	}

	for _, menu := range menus {
		if menu.ParentID != nil && menuMap[*menu.ParentID] != nil {
			parentMenu := menuMap[*menu.ParentID]
			if children, ok := (*parentMenu)["children"].([]map[string]interface{}); ok {
				(*parentMenu)["children"] = append(children, *menuMap[menu.ID])
			}
		}
	}

	for _, menu := range menus {
		if menu.ParentID == nil {
			transformedMenus = append(transformedMenus, *menuMap[menu.ID])
		}
	}

	return transformedMenus
}
