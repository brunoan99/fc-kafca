package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	err := CreateTopic()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Topic criado com sucesso")
}

func CreateTopic() error {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	admin, err := kafka.NewAdminClient(configMap)
	if err != nil {
		return err
	}
	topic := kafka.TopicSpecification{
		Topic:             "teste",
		NumPartitions:     3,
		ReplicationFactor: 1,
	}
	topics := []kafka.TopicSpecification{topic}
	_, err = admin.CreateTopics(context.Background(), topics)
	return err
}
