package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/fire_gateway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))

	// 自动迁移模型，确保模型与数据库表结构保持同步
	err = DB.AutoMigrate(
		&User{},
		&Role{},
		&JwtToken{},
		&UserRole{},
		&APIEndpoint{},
		&Route{},
		&Menu{},
	)
	if err != nil {
		return err
	}

	return nil
}
