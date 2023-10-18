package dao

import "gateway/models"

var DefaultProjectDAO *ProjectDAO

var DefaultProjectSettingDAO *ProjectSettingDAO

func InitDao() {
	DefaultProjectDAO = NewProjectDAO(models.DB)
	DefaultProjectSettingDAO = NewProjectSettingDAO(models.DB)
}
