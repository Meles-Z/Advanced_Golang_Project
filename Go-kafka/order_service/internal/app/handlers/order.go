package handers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
	"github.com/meles-z/go-kafa/order_service/internal/app/service"
)

type OrderHandler struct {
	UseCase *service.OrderUseCase
}

func NewOrderHandler(useCase *service.OrderUseCase) *OrderHandler {
	return &OrderHandler{UseCase: useCase}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order entities.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.Status = "pending"

	if err := h.UseCase.CreateOrder(&order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
