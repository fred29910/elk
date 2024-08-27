package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// const ntsUrl = "127.0.0.1:4222"

func main() {

	// 连接到本地 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer nc.Close()

	// 要发布的主题
	subject := "updates"

	// 要发布的数据
	data := "Hello, NATS!"

	// 发布消息
	err = nc.Publish(subject, []byte(data))
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Published data '%s' to subject '%s'\n", data, subject)

	// 刷新队列，确保消息已发送
	nc.Flush()

	// 检查是否有错误
	if err := nc.LastError(); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Data pushed successfully!")
}
