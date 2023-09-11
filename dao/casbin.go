package dao

import (
	"gateway/models"
	"gorm.io/gorm"
)

type CasbinRuleDAO struct {
	DB *gorm.DB
}

func (dao *CasbinRuleDAO) GetPolicies() ([]models.CasbinRule, error) {
	var rules []models.CasbinRule
	err := dao.DB.Find(&rules).Error
	return rules, err
}

func (dao *CasbinRuleDAO) AddPolicy(rule models.CasbinRule) error {
	return dao.DB.Create(&rule).Error
}
