package utils

var Id *AutoInc

func init() {
	Id = newGenderID(100000, 1)
}

type AutoInc struct {
	start   int
	step    int
	running bool
	queue   chan int
}

// newGenderID 创建一个id生成器
func newGenderID(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

// Next 获取一个id
func (ai *AutoInc) Next() int {
	return <-ai.queue
}

// Close 关闭生成器
func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}
