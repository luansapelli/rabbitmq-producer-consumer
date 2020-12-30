package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/luansapelli/rabbitmq-producer-consumer/config"
	"github.com/luansapelli/rabbitmq-producer-consumer/utils"
	"github.com/streadway/amqp"
	"log"
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
	InitRabbitProducer()
}

func InitRabbitProducer() {

	rabbitConfig := config.RabbitMqConfig()

	defer rabbitConfig.Conn.Close()

	// Generate UUID for each message
	UUID, _ := uuid.NewRandom()
	messageUuid := UUID.String()

	// Message to RabbitMQ
	message := RabbitMessage{
		Uuid:        messageUuid,
		ServiceName: "rabbitmq-producer-consumer",
		Name:        "Luan Sapelli",
		Age:         21,
		Country:     "Brazil",
		Message:     "Hello World!",
	}

	body, err := json.Marshal(message)
	if err != nil {
		utils.HandleError(err, "Error encoding JSON")
	}

	err = rabbitConfig.Channel.Publish("", rabbitConfig.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("Message sent: %v", message)
}
