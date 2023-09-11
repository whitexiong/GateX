package handlers

import (
	"gateway/dao"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllRoutes(c *gin.Context) {
	routes := dao.GetAllRoutes()
	transformedRoutes := ConvertRoutesToTree(routes)
	SendResponse(c, http.StatusOK, 200, transformedRoutes)
	return
}

func GetRoute(c *gin.Context) {
	var route models.Route
	id := c.Param("id")
	models.DB.Find(&route, id)
	SendResponse(c, http.StatusOK, 200, route)
	return
}

func ConvertRoutesToTree(routes []models.Route) []map[string]interface{} {
	var transformedRoutes []map[string]interface{}
	routeMap := make(map[uint]*map[string]interface{})

	for _, route := range routes {
		transformedRoute := map[string]interface{}{
			"id":       route.ID,
			"value":    route.ID,
			"name":     route.Name,
			"label":    route.Name,
			"path":     route.Path,
			"status":   route.Status,
			"children": []map[string]interface{}{},
		}
		routeMap[route.ID] = &transformedRoute
	}

	for _, route := range routes {
		if route.ParentID != nil && routeMap[*route.ParentID] != nil {
			parentRoute := routeMap[*route.ParentID]
			if children, ok := (*parentRoute)["children"].([]map[string]interface{}); ok {
				(*parentRoute)["children"] = append(children, *routeMap[route.ID])
			}
		}
	}

	for _, route := range routes {
		if route.ParentID == nil {
			transformedRoutes = append(transformedRoutes, *routeMap[route.ID])
		}
	}

	return transformedRoutes
}

func CreateRoute(c *gin.Context) {
	var route models.Route

	if err := c.BindJSON(&route); err != nil {
		c.Error(err)
		return
	}

	if result := models.DB.Create(&route); result.Error != nil {
		c.Error(result.Error)
		return
	}

	SendResponse(c, http.StatusOK, 200, route)
}

func GetRouteByID(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"apierrors": "Route entry not found"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func UpdateRoute(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.Error(result.Error)
		return
	}

	if err := c.BindJSON(&route); err != nil {
		c.Error(err)
		return
	}

	if result := models.DB.Save(&route); result.Error != nil {
		c.Error(result.Error)
		return
	}

	SendResponse(c, http.StatusOK, 200, route)
}

func DeleteRoute(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.Error(result.Error)
		return
	}

	if result := models.DB.Delete(&route, id); result.Error != nil {
		c.Error(result.Error)
		return
	}

	SendResponse(c, http.StatusOK, 200, nil)
}

func GetRoutePathList(c *gin.Context) {
	queryString := c.DefaultQuery("path", "")

	var routes []models.Route
	result := models.DB.Where("Path LIKE ?", "%"+queryString+"%").Find(&routes)
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	paths := make([]map[string]string, len(routes))
	for i, route := range routes {
		paths[i] = map[string]string{
			"id":    strconv.Itoa(int(route.ID)),
			"name":  route.Name,
			"value": route.Path,
		}
	}

	SendResponse(c, http.StatusOK, 200, paths)
	return
}
