package dao

import (
	"errors"
	"gateway/models"
	"gorm.io/gorm"
)

type ProjectAPIEndpointDAO struct {
	DB *gorm.DB
}

func NewProjectAPIEndpointDAO(db *gorm.DB) *ProjectAPIEndpointDAO {
	return &ProjectAPIEndpointDAO{DB: db}
}

func (dao *ProjectAPIEndpointDAO) CreateEndpoint(e *models.ProjectAPIEndpoint) error {
	if e == nil {
		return errors.New("provided endpoint is nil")
	}
	return dao.DB.Create(e).Error
}

func (dao *ProjectAPIEndpointDAO) GetEndpointByID(id uint) (*models.ProjectAPIEndpoint, error) {
	var endpoint models.ProjectAPIEndpoint
	err := dao.DB.First(&endpoint, id).Error
	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}

func (dao *ProjectAPIEndpointDAO) UpdateEndpoint(e *models.ProjectAPIEndpoint) error {
	if e == nil {
		return errors.New("provided endpoint is nil")
	}
	return dao.DB.Save(e).Error
}

func (dao *ProjectAPIEndpointDAO) DeleteEndpoint(id uint) error {
	return dao.DB.Delete(&models.ProjectAPIEndpoint{}, id).Error
}

func (dao *ProjectAPIEndpointDAO) ListEndpointsByProject(projectID uint) ([]models.ProjectAPIEndpoint, error) {
	var endpoints []models.ProjectAPIEndpoint
	err := dao.DB.Where("project_id = ?", projectID).Find(&endpoints).Error
	return endpoints, err
}
