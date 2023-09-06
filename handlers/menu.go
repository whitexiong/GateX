package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取所有菜单
func GetAllMenus(c *gin.Context) {
	var menus []models.Menu
	models.DB.Order("`Order`").Find(&menus)

	transformedMenus := ConvertMenusToTree(menus)
	SendResponse(c, http.StatusOK, 200, transformedMenus, "Success")
}

// 创建新菜单
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Bad request data")
		return
	}
	if result := models.DB.Create(&menu); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to create menu")
		return
	}
	SendResponse(c, http.StatusOK, 200, menu, "Success")
}

// 获取指定ID的菜单
func GetMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		SendResponse(c, http.StatusNotFound, 404, nil, "Menu not found")
		return
	}
	SendResponse(c, http.StatusOK, 200, menu, "Success")
}

// 更新指定ID的菜单
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		SendResponse(c, http.StatusNotFound, 404, nil, "Menu not found")
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Bad request data")
		return
	}
	if result := models.DB.Save(&menu); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to update menu")
		return
	}
	SendResponse(c, http.StatusOK, 200, menu, "Success")
}

// 删除指定ID的菜单
func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		SendResponse(c, http.StatusNotFound, 404, nil, "Menu not found")
		return
	}
	if result := models.DB.Delete(&menu); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to delete menu")
		return
	}
	SendResponse(c, http.StatusOK, 200, nil, "Menu deleted successfully")
}

func ConvertMenusToTree(menus []models.Menu) []map[string]interface{} {
	var transformedMenus []map[string]interface{}
	menuMap := make(map[uint]*map[string]interface{})

	// 创建基础格式
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
