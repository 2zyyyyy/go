package main

import (
	"fmt"
	"go/github.io/2zyyyyy/chineseDocumentation/MQ/rabbitMQ"
)

func main() {
	rabbitMQ := rabbitMQ.NewRabbitMQSimple("" + "simple")
	rabbitMQ.PublishSimple("Hello MQ!")
	fmt.Println("发送成功~")
}
