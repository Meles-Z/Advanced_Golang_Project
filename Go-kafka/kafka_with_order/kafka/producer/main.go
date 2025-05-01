package main

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	// kafka broker address
	broker = "localhost:9092"
)

type Message struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	// create a new kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   broker,
		"acks":                "all", // Ensure all replicas acknowledge
		"go.delivery.reports": true,  // Enable delivery reports
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	// delivery report handler
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("Message delivered to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	// create a new message
	message := []Message{
		{
			Name:  "Meles Zawude",
			Age:   30,
			Email: "meles@gmail.com",
			Phone: "1234567890",
		},
		{
			Name:  "Meseret Zawude",
			Age:   45,
			Email: "meseret@gmail.com",
			Phone: "1234567890",
		},
		{
			Name:  "FFFF Zawude",
			Age:   45,
			Email: "meseret@gmail.com",
			Phone: "1234567890",
		},
		{
			Name:  "Langanoo Zawude",
			Age:   45,
			Email: "lanret@gmail.com",
			Phone: "1234567890",
		},
		{
			Name:  "Langanooajgaskgjas Zawude",
			Age:   45,
			Email: "lanret@gmail.com",
			Phone: "1234567890",
		},
	}

	msg, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}
	var topic = "test-topic"
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, nil)
	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
	}
	producer.Flush(5000) // Wait for message delivery
	log.Printf("Produced message: %s\n", msg)

}
