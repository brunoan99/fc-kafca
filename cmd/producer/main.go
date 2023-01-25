package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer := NewKafkaProducer()
	Publish("mensagem", "teste", producer, nil)
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMag := &kafka.ConfigMap{
		"bootstrap.servers": "kafka-kafka-1:9092",
	}
	producer, err := kafka.NewProducer(configMag)
	if err != nil {
		log.Println(err.Error())
	}
	return producer
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, nil)
	return err
}
