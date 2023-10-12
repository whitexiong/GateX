package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabase() error {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("加载.env文件错误: %s", err)
	}

	// 从环境变量中获取数据库连接信息
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	// 构建DSN字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	// 自动迁移模型，确保模型与数据库表结构保持同步
	err = DB.AutoMigrate(
		&User{},
		&Role{},
		&JwtToken{},
		&UserRole{},
		&APIEndpoint{},
		&Route{},
		&Menu{},
		&ChatRoom{},
		&Message{},
		&ChatRoomUser{},
		&Project{},
		&ProjectAPIEndpoint{},
	)
	if err != nil {
		return err
	}

	return nil
}
