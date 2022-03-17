package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_topic"
	msg.Value = sarama.StringEncoder("this is my test message!")

	// 连接 kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("connection kafka failed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("pid: %v, offset:= %v\n", pid, offset)
}
