package mq

import "log"

var done chan bool

// 开始监听队列，获取消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {

	// 获取消息信道
	msgs, err := channel.Consume(qName, cName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	done = make(chan bool)
	// 循环获取队列的消息,为了防止循环一直阻塞代码，用goroutine包裹起来
	go func() {
		for msg := range msgs {
			processSuc := callback(msg.Body)
			if !processSuc {
				//TODO：没有执行成功，写到另一个队列，用于异常情况的重试

			}
		}
	}()
	// done没有新的消息，则会一直阻塞下去
	<-done
	// 当有消息过来，也就说明要关闭消费者了
	channel.Close()
}
