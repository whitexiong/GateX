package dao

import "gateway/models"

var DefaultProjectDAO *ProjectDAO

func InitDao() {
	DefaultProjectDAO = NewProjectDAO(models.DB)
}
