package dao

import (
	"errors"
	"gateway/models"
	"gorm.io/gorm"
)

type ProjectSettingDAO struct {
	DB *gorm.DB
}

func NewProjectSettingDAO(db *gorm.DB) *ProjectSettingDAO {
	return &ProjectSettingDAO{DB: db}
}

func (dao *ProjectSettingDAO) GetProjectSettingByProjectID(projectID uint) (*models.ProjectSetting, error) {
	var setting models.ProjectSetting
	err := dao.DB.Where("project_id = ?", projectID).First(&setting).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (dao *ProjectSettingDAO) UpdateProjectSetting(p *models.ProjectSetting) error {
	return dao.DB.Save(p).Error
}

func (dao *ProjectSettingDAO) CreateProjectSetting(ps *models.ProjectSetting) error {
	if ps == nil {
		return errors.New("provided project setting is nil")
	}
	return dao.DB.Create(ps).Error
}
