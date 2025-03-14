package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumeOrders(broker, topic string, useCase *usecase.PaymentUseCase) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "payment_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var order domain.Order
			json.Unmarshal(msg.Value, &order)

			log.Println("Processing payment for order:", order.ID)
			useCase.ProcessPayment(order)
		} else {
			log.Println("Consumer error:", err)
		}
	}
}
