package dao

import (
	"errors"
	"gateway/models"
	"gorm.io/gorm"
)

type ProjectDAO struct {
	DB *gorm.DB
}

func NewProjectDAO(db *gorm.DB) *ProjectDAO {
	return &ProjectDAO{DB: db}
}

func (dao *ProjectDAO) CreateProject(p *models.Project) error {
	if p == nil {
		return errors.New("provided projects is nil")
	}
	return dao.DB.Create(p).Error
}

func (dao *ProjectDAO) GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	err := dao.DB.Preload("APIs").First(&project, id).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (dao *ProjectDAO) UpdateProject(p *models.Project) error {
	if p == nil {
		return errors.New("provided project is nil")
	}
	return dao.DB.Model(p).Updates(p).Error
}

func (dao *ProjectDAO) DeleteProject(id uint) error {
	return dao.DB.Delete(&models.Project{}, id).Error
}

func (dao *ProjectDAO) ListProjects() ([]models.Project, error) {
	var projects []models.Project
	err := dao.DB.Find(&projects).Error
	return projects, err
}
