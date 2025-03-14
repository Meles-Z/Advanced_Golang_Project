package service

import (
	"log"

	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
	"github.com/meles-z/go-kafa/order_service/internal/app/kafka"
	repository "github.com/meles-z/go-kafa/order_service/internal/app/respository"
)

type OrderUseCase struct {
	Repo     *repository.OrderRepository
	Producer *kafka.KafkaProducer
}

func NewOrderUseCase(repo *repository.OrderRepository, producer *kafka.KafkaProducer) *OrderUseCase {
	return &OrderUseCase{Repo: repo, Producer: producer}
}

func (uc *OrderUseCase) CreateOrder(order *entities.Order) error {
	if err := uc.Repo.Save(order); err != nil {
		return err
	}
	log.Println("Order saved, publishing event to Kafka...")
	return uc.Producer.PublishOrder(order)
}
