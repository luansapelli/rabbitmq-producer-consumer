package main

import (
	"encoding/json"
	"github.com/luansapelli/rabbitmq-producer-consumer/config"
	"github.com/luansapelli/rabbitmq-producer-consumer/utils"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main (){
	InitRabbitConsumer()
}

func InitRabbitConsumer() {

	var messageChannel <-chan amqp.Delivery
	var messageJSON map[string]interface{}

	rabbitConfig := config.RabbitMqConfig()

	defer rabbitConfig.Conn.Close()

	messageChannel, rabbitConfig.RbmqErr = rabbitConfig.Channel.Consume(
		rabbitConfig.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.HandleError(rabbitConfig.RbmqErr, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())

		for i := range messageChannel {
			log.Printf("Message received: %s", i.Body)

			err := json.Unmarshal(i.Body, &messageJSON)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			if err := i.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	// Stop for program termination
	<-stopChan

}
