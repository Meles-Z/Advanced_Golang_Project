package repository

import (
	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order *entities.Order) error {
	err := r.DB.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}
