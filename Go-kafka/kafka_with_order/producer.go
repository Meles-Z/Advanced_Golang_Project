package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	topic := "orders"
	message := []string{
		"order-3: Buy 100 AAPL",
		"order-4: Sell 50 TSLA",
		"order-5: Buy 10 BTC",
	}

	for _, msg := range message {
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			log.Printf("Failed to produce message: %v\n", err)
		} else {
			log.Printf("Produced message: %s\n", msg)
		}
		producer.Flush(5000)
		fmt.Println("Message sent to Kafka topic:", topic)
	}
}
