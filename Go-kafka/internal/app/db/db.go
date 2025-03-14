package db

import (
	"fmt"

	"github.com/meles-z/go-kafa/configs"
	"github.com/meles-z/go-kafa/internal/app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d name=%s username=%s password=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.Username, cfg.Password)
	normalDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = normalDb
	normalDb.AutoMigrate(&entities.User{})
	return normalDb, nil
}
