package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/luansapelli/rabbitmq-producer-consumer/config"
	"github.com/luansapelli/rabbitmq-producer-consumer/helper"
	"github.com/streadway/amqp"
)

func main() {
	rabbit := config.RabbitConfig()
	startConsumer(rabbit)
}

func startConsumer(rabbit *config.RabbitMQ) {
	var channel <-chan amqp.Delivery
	var messageInterface map[string]interface{}

	defer rabbit.Connection.Close()

	channel, rabbit.Error = rabbit.Channel.Consume(
		rabbit.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	helper.FailOnError(rabbit.Error, "could not register consumer")

	stopChannel := make(chan bool)

	go func() {
		log.Printf("consumer ready, PID: %v", os.Getpid())

		for message := range channel {
			log.Printf("message received: %s", message.Body)

			err := json.Unmarshal(message.Body, &messageInterface)
			if err != nil {
				log.Printf("error decoding Json: %s", err)
			}

			if err := message.Ack(false); err != nil {
				log.Printf("error acknowledging message : %s", err)
			} else {
				log.Printf("acknowledged message")
			}
		}
	}()

	// Stop for program termination
	<-stopChannel
}
