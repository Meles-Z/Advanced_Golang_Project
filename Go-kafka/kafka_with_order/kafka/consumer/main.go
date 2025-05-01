package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker = "localhost:9092"
	topic  = "test-topic"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "test-group",
		"auto.offset.reset": "latest", // Start reading at the latest message
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()
	// Subscribe to the topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}
	// Start consuming messages
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %s", err)
			continue
		}
		log.Printf("Received message: %s\n", string(msg.Value))
	}

}
