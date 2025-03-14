package db

import (
	"fmt"

	"github.com/meles-z/go-kafa/order_service/configs"
	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password)
	normalDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = normalDb
	normalDb.AutoMigrate(&entities.User{})
	return normalDb, nil
}
