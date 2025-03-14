package repository

import (
	"database/sql"

	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *entities.Order) error {
	_, err := r.DB.Exec("INSERT INTO orders (id, amount, status, created_at) VALUES ($1, $2, $3, $4)",
		order.ID, order.Amount, order.Status, order.CreatedAt)
	return err
}
