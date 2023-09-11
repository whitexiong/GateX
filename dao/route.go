package dao

import (
	"gateway/models"
)

func GetAllRoutes() []models.Route {
	var routes []models.Route
	models.DB.Order("created_at").Find(&routes)
	return routes
}
