package mq

// TransferData : 将要写到rabbitmq的数据的结构体
type TransferData struct {
	TaskID int64
	Data   interface{}
	msg    string
}
