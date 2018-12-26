package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pibigstar/go-todo/cron/jobs"
	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/robfig/cron"
)

var log = logger.New("cron")

func main() {
	c := cron.New()
	for _, job := range jobs.GetJobs() {
		log.Info("job启动", "job name", job.Name())
		c.AddFunc(job.Cron(), func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("job 运行出错", "job name", job.Name(), "error", err)
				}
			}()
			// 执行任务
			job.Run()
		})
	}
	c.Start()
	defer c.Stop()

	// 阻塞main协程退出，直到手动退出程序
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	killSignal := <-interrupt

	log.Info("退出定时任务", "signal", killSignal)
}
