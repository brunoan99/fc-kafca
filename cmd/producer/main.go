package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("mensagem", "teste", producer, []byte("teste"), deliveryChan)
	go DeliveryReport(deliveryChan)
	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	kafkaPath := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_PATH"), os.Getenv("KAFKA_PORT"))
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   kafkaPath,
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}
	producer, err := kafka.NewProducer(configMap)
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

func DeliveryReport(deliveryChan chan kafka.Event) {
	for event := range deliveryChan {
		switch ev := event.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada: ", ev.Value)
				fmt.Println("Mensagem armazenada em: ", ev.TopicPartition)
			}
		}
	}
}
