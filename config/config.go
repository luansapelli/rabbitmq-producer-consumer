package config

import (
	"github.com/luansapelli/rabbitmq-producer-consumer/helper"
	goRabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Queue      amqp.Queue
	Channel    *amqp.Channel
	Connection *amqp.Connection
	Error      error
}

func RabbitConfig() *RabbitMQ {
	rabbit := new(RabbitMQ)

	rabbit.Connection, rabbit.Error = amqp.Dial(goRabbit.Config.AMQPConnectionURL)
	helper.FailOnError(rabbit.Error, "failed to connect on RabbitMQ")

	rabbit.Channel, rabbit.Error = rabbit.Connection.Channel()
	helper.FailOnError(rabbit.Error, "failed to open a channel")

	rabbit.Queue, rabbit.Error = rabbit.Channel.QueueDeclare(
		"myqueue",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	helper.FailOnError(rabbit.Error, "could not declare `myqueue` queue")

	err := rabbit.Channel.Qos(1, 0, false)
	helper.FailOnError(err, "could not configure QoS")

	return rabbit
}
