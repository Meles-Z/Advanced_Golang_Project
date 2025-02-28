package db

import (
	"fmt"

	"github.com/meles-z/edusysnc_grpc/configs"
	"github.com/meles-z/edusysnc_grpc/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *configs.DatabaseConfigruration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s name=%s port=%d user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Name, cfg.Port, cfg.Username, cfg.Password)
	studentDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db:%s", err)
	}

	DB = studentDb
	err = studentDb.AutoMigrate(
		&models.Address{},
		&models.Course{},
		&models.Department{},
		&models.Student{},
		&models.Teacher{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to automigrate database:%s", err)
	}
	fmt.Println("Database is connected successfully!")
	return studentDb, nil
}
