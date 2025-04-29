package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Order struct {
	ID     string `json:"id"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

var producer *kafka.Producer

func main() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	http.HandleFunc("/order", handleOrder)
	log.Println("Order service listening on :8000")
	http.ListenAndServe(":8000", nil)
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	value, _ := json.Marshal(order)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &[]string{"orders"}[0], Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	w.WriteHeader(http.StatusAccepted)
}
