package kafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/meles-z/go-kafa/order_service/internal/app/entities"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatal(err)
	}
	return &KafkaProducer{Producer: p, Topic: topic}
}

func (kp *KafkaProducer) PublishOrder(order *entities.Order) error {
	orderJSON, _ := json.Marshal(order)
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny},
		Value:          orderJSON,
	}, nil)
	return err
}
