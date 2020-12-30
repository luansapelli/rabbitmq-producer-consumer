# Golang RabbitMQ Producer/Consumer 🚀

<p>
This service send a message to RabbitMQ queue and consume this message.
</p>

### Tools

* [RabbitMQ](https://www.rabbitmq.com/)

### Step by step

1. Start RabbitMQ using the `docker-compose-yaml` file (make sure the Docker is running)
2. Access `http://localhost:15672/` to make sure if RabbitMQ is running (username: `guest`, password: `guest`)
3. Execute `go run consumer/consumer.go`, in terminal you can see this message: `Consumer ready, PID: 55725` (let it run in a separated terminal)
4. Access `http://localhost:15672/#/channels` and you can see an open channel, that is your consumer
5. Execute `go run producer/producer.go`, to send a message (you can send more than one message, and you can edit this message in `producer.go`). The message will be printed in terminal, something like (`Message sent: {ecbe7cd6-ecd1-4079-b1d0-796790a440cc rabbitmq-producer-consumer Luan Sapelli 21 Brazil Hello World!}`)
6. In the consumer terminal you will see this message: `Message received: {"uuid":"6cd8782f-4635-4773-b153-f53789ba2bcb","serviceName":"rabbitmq-producer-consumer","name":"Luan Sapelli","age":21,"country":"Brazil","message":"Hello World!"}`
