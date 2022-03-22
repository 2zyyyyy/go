package main

import "go/github.io/2zyyyyy/chineseDocumentation/MQ/rabbitMQ"

func main() {
	rabbitmq := rabbitMQ.NewRabbitMQSimple("" + "simpleU")
	rabbitmq.ConsumerSimple()
}
