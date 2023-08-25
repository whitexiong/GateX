package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取所有菜单
func GetAllMenus(c *gin.Context) {
	// 获取数据库中的所有菜单数据
	var menus []models.Menu
	models.DB.Order("`Order`").Find(&menus) // 可能需要根据Order字段排序

	// 转换为新格式
	transformedMenus := ConvertMenusToTree(menus)

	c.JSON(200, gin.H{
		"menus": transformedMenus,
	})
}

func ConvertMenusToTree(menus []models.Menu) []map[string]interface{} {
	var transformedMenus []map[string]interface{}
	menuMap := make(map[uint]*map[string]interface{})

	// 创建基础格式
	for _, menu := range menus {
		transformedMenu := map[string]interface{}{
			"id":       menu.ID,
			"name":     menu.Name,
			"date":     menu.CreatedAt.Format("2006-01-02"),
			"path":     menu.Path,
			"icon":     menu.Icon,
			"children": []map[string]interface{}{},
		}
		menuMap[menu.ID] = &transformedMenu
	}

	// 添加子菜单到对应的父菜单
	for _, menu := range menus {
		if menu.ParentID != nil && menuMap[*menu.ParentID] != nil {
			parentMenu := menuMap[*menu.ParentID]
			if children, ok := (*parentMenu)["children"].([]map[string]interface{}); ok {
				(*parentMenu)["children"] = append(children, *menuMap[menu.ID])
			}
		}
	}

	// 只选择顶级菜单(没有父ID的菜单)来构建最终的列表
	for _, menu := range menus {
		if menu.ParentID == nil {
			transformedMenus = append(transformedMenus, *menuMap[menu.ID])
		}
	}

	// 添加hasChildren字段
	for _, menu := range transformedMenus {
		children := menu["children"].([]map[string]interface{})
		if len(children) > 0 {
			menu["hasChildren"] = true
		}
	}

	return transformedMenus
}

// 创建新菜单
func CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}
	if result := models.DB.Create(&menu); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// 获取指定ID的菜单
func GetMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// 更新指定ID的菜单
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}
	if result := models.DB.Save(&menu); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu"})
		return
	}
	c.JSON(http.StatusOK, menu)
}

// 删除指定ID的菜单
func DeleteMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id")
	if result := models.DB.First(&menu, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}
	if result := models.DB.Delete(&menu); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
