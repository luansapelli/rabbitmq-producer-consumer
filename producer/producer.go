package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/luansapelli/rabbitmq-producer-consumer/config"
	"github.com/luansapelli/rabbitmq-producer-consumer/helper"
	"github.com/streadway/amqp"
)

type RabbitMessage struct {
	Uuid        string `json:"uuid"`
	ServiceName string `json:"serviceName"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Country     string `json:"country"`
	Message     string `json:"message"`
}

func main() {
	rabbit := config.RabbitConfig()
	startProducer(rabbit)
}

func startProducer(rabbit *config.RabbitMQ) {
	defer rabbit.Connection.Close()

	// Generate UUID for each message
	uuid := uuid.NewString()

	// Message to RabbitMQ
	message := RabbitMessage{
		Uuid:        uuid,
		ServiceName: "rabbitmq-producer-consumer",
		Name:        "Luan Felipe Sapelli",
		Age:         21,
		Country:     "Brazil",
		Message:     "Hello World!",
	}

	rawMessage, err := json.Marshal(message)
	if err != nil {
		helper.FailOnError(err, "error to marshal message")
	}

	err = rabbit.Channel.Publish("", rabbit.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         rawMessage,
	})
	if err != nil {
		helper.FailOnError(err, "error to publish message")
	}

	log.Printf("message sent: %v", message)
}
