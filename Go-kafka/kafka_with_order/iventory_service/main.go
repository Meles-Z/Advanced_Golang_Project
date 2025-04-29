package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Order struct {
	ID     string `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "inventory-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"orders"}, nil)

	log.Println("Inventory service is consuming from Kafka...")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var order Order
			json.Unmarshal(msg.Value, &order)
			log.Printf("Received order: %+v. Deducting inventory...\n", order)
			// TODO: Update PostgreSQL inventory table.
		} else {
			log.Printf("Error reading message: %v\n", err)
		}
	}
}
