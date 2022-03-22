package rabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://simpleU:123456@127.0.0.1:5672/simple"

// rabbitMQ 结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机名称
	Key       string // bind key 名称
	Mqurl     string // 连接信息
}

// 结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// 关闭 channel 和 connection
func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// 错误处理函数
func (*RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", msg, err)
		panic("error:" + msg)
	}
}

// 创建简单模式下的 RabbitMQ
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 创建 RabbitMQ 实例
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取 connection
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnError(err, "failed to connect rabbitmq!")
	// 获取 channel
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnError(err, "failed to open channel!")
	return rabbitMQ
}

// 直接模式队列生产
func (r *RabbitMQ) PublishSimple(msg string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞处理
		nil,   //额外的属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 调用 channel，发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false, //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

// simple 模式下的消费者
func (r *RabbitMQ) ConsumerSimple() {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞处理
		nil,   //额外的属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for m := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", m.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
