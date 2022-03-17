package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka 消费者
func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("test_topic") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return
	}
	fmt.Println(partitionList)
	// 遍历所有分区
	for partition := range partitionList {
		// 针对每个分区 创建一个分区对应的消费者
		pc, err := consumer.ConsumePartition("test_topic", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
