package config

import (
	"github.com/luansapelli/rabbitmq-producer-consumer/utils"
	golangrabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

type RbmqConfig struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
	RbmqErr error
}

func RabbitMqConfig() *RbmqConfig {

	config := &RbmqConfig{}

	config.Conn, config.RbmqErr = amqp.Dial(golangrabbit.Config.AMQPConnectionURL)
	utils.HandleError(config.RbmqErr, "Failed to connect to RabbitMQ")

	config.Channel, config.RbmqErr = config.Conn.Channel()
	utils.HandleError(config.RbmqErr, "Failed to open a channel")

	config.Queue, config.RbmqErr = config.Channel.QueueDeclare(
		"myqueue",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.HandleError(config.RbmqErr, "Could not declare `myqueue` queue")

	RbmqErr := config.Channel.Qos(1, 0, false)
	utils.HandleError(RbmqErr, "Could not configure QoS")

	return config
}
