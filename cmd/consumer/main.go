package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer := newKafkaConsumer()
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		} else {
			log.Println(err.Error())
		}
	}
}

func newKafkaConsumer() *kafka.Consumer {
	kafkaPath := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_PATH"), os.Getenv("KAFKA_PORT"))
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaPath,
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
	}
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)
	return consumer
}
