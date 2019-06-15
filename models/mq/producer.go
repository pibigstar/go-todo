package mq

import (
	"fmt"
	"log"

	"github.com/pibigstar/go-todo/constant"
	"github.com/streadway/amqp"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

func initChannel() bool {
	// 判断channel是否已经被初始化了
	if channel != nil {
		return true
	}
	// 获取rabbitmq的一个连接
	conn, err := amqp.Dial(constant.RabbitURL)
	if err != nil {
		fmt.Printf("Failed to dial the rabbitmq,err:%s\n", err.Error())
		return false
	}
	// 打开一个channel，用户消息的发布与接收
	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// 发布消息
func Publish(exchange, routingkey string, msg []byte) bool {
	// 判断channel是否正常
	if !initChannel() {
		return false
	}
	pubMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}
	err := channel.Publish(exchange, routingkey, false, false, pubMsg)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func init() {
	initChannel()
}
