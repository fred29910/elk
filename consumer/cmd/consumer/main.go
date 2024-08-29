package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 连接到 RabbitMQ 服务器
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const (
	mqURL        = "amqp://dev:dev@localhost:5672/"
	exchangeName = "amq_tracking"
	exchangeType = "direct"
	routingKey   = "tracking.key"

	queueName = "tracking.data"
)

func main() {
	conn, err := amqp.Dial(mqURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		true,      // 是否持久化
		false,     // 是否自动删除
		false,     // 是否排他
		false,     // 是否阻塞
		nil,       // 其他参数
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	err = ch.QueueBind(
		q.Name,       // name of the queue
		routingKey,   // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	// 发送消息
	body := "Hello World!"
	err = ch.Publish(
		exchangeName, // 交换机名称
		// q.Name,       // 路由键，即队列名称
		routingKey, // 路由键，即队列名称
		false,      // 是否强制
		false,      // 是否立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Printf(" [x] Sent %s\n", body)

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者名称
		true,   // 自动应答
		false,  // 是否排他
		false,  // 是否阻塞
		false,  // 其他参数
		nil,    // 其他参数
	)
	failOnError(err, "Failed to register a consumer")

	// 使用 Goroutine 来处理消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	select {}
}
