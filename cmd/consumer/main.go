package main

import (
	"fmt"
	"log"

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
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-kafka-1:9092",
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
