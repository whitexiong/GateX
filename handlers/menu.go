package handlers

import (
	"fmt"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetUserMenus(c *gin.Context) {
	// 获取用户名从 JWT 中
	roleName, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in token"})
		return
	}

	// 如果是 root 用户，直接返回所有菜单项
	if roleName == "root" {
		var allMenus []models.Menu
		if err := models.DB.Find(&allMenus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all menus for root"})
			return
		}
		transformedMenus := ConvertMenusToTree(allMenus)
		SendResponse(c, http.StatusOK, 200, transformedMenus, "Success")
		return
	}

	// 获取用户ID从 JWT 中
	role, _ := c.Get("role")
	if role == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in token"})
		return
	}

	// 使用用户ID获取菜单
	menus, err := fetchUserMenus(fmt.Sprintf("%v", role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve menus"})
		return
	}

	// 转换菜单列表为树形结构
	transformedMenus := ConvertMenusToTree(menus)
	SendResponse(c, http.StatusOK, 200, transformedMenus, "Success")
}

func fetchUserMenus(roleName string) ([]models.Menu, error) {
	// 1. 通过用户ID查询用户的所有角色
	var casbinRules []models.CasbinRule
	if err := models.DB.Where("v0 = ?", roleName).Find(&casbinRules).Error; err != nil {
		return nil, err
	}

	var routeIds []uint
	for _, cr := range casbinRules {
		routeId, _ := strconv.ParseUint(cr.V3, 10, 32)
		routeIds = append(routeIds, uint(routeId))
	}

	// 3. 根据路由ID从Menu表中查询出用户具有权限访问的所有菜单项
	var menus []models.Menu
	if err := models.DB.Where("route_id IN ?", routeIds).Find(&menus).Error; err != nil {
		return nil, err
	}

	fmt.Println("routids:", routeIds)
	fmt.Println("menus:", menus)
	return menus, nil
}
