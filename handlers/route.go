package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRoutes(c *gin.Context) {
	var routes []models.Route
	models.DB.Order("created_at").Find(&routes)

	transformedRoutes := ConvertRoutesToTree(routes)

	c.JSON(200, gin.H{
		"routes": transformedRoutes,
	})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if result := models.DB.Create(&route); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route entry"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func GetRouteByID(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route entry not found"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func UpdateRouteByID(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route entry not found"})
		return
	}

	if err := c.BindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if result := models.DB.Save(&route); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update route entry"})
		return
	}

	c.JSON(http.StatusOK, route)
}

func DeleteRouteByID(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route entry not found"})
		return
	}

	if result := models.DB.Delete(&route); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete route entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Route entry deleted successfully"})
}
