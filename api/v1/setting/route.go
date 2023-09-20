package setting

import (
	"gateway/apierrors"
	"gateway/dao"
	"gateway/models"
	"gateway/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllRoutes(c *gin.Context) {
	routes := dao.GetAllRoutes()
	transformedRoutes := util.ConvertToTree(routes, util.MapRouteToTreeItem)
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

func CreateRoute(c *gin.Context) {
	var route models.Route

	if err := c.BindJSON(&route); err != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	if result := models.DB.Create(&route); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	SendResponse(c, http.StatusOK, 200, route)
}

func UpdateRoute(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	if err := c.BindJSON(&route); err != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	if result := models.DB.Save(&route); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	SendResponse(c, http.StatusOK, 200, route)
}

func DeleteRoute(c *gin.Context) {
	var route models.Route
	id := c.Param("id")

	if result := models.DB.First(&route, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	if result := models.DB.Delete(&route, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
		return
	}

	SendResponse(c, http.StatusOK, 200, nil)
}

func GetRoutePathList(c *gin.Context) {
	queryString := c.DefaultQuery("path", "")

	var routes []models.Route
	result := models.DB.Where("Path LIKE ?", "%"+queryString+"%").Find(&routes)
	if result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DatabaseError, nil)
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
