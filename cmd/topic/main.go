package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	kafkaPath := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_PATH"), os.Getenv("KAFKA_PORT"))
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaPath,
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
