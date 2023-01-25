package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("mensagem", "teste", producer, nil, deliveryChan)

	event := <-deliveryChan
	msg := event.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar")
	} else {
		fmt.Println("Mensagem enviada: ", msg.Value)
		fmt.Println("Mensagem armazenada em:", msg.TopicPartition)
	}

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

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	return err
}
