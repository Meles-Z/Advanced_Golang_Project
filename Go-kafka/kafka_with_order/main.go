package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "eventory-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"orders"}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	fmt.Println("Kafka consumer started...")

	for {
		msg, err := consumer.ReadMessage(100 * time.Millisecond)
		if err != nil {
			// Handle timeout or other recoverable errors
			kErr, ok := err.(kafka.Error)
			if ok && kErr.Code() == kafka.ErrTimedOut {
				continue
			}
			log.Printf("Consumer error: %v\n", err)
			continue
		}
		fmt.Println("Message received:", string(msg.Value))
		fmt.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
