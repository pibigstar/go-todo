package jobs

import (
	"fmt"
	"time"
)

func init() {
	remind := &RemindRunner{}
	addJobs(remind)
}

// RemindRunner 任务提醒
type RemindRunner struct {
}

// Name 任务名
func (*RemindRunner) Name() string {
	return "RemindRunner"
}

// Cron 任务执行间隔 每分钟执行一次
func (*RemindRunner) Cron() string {
	return "0 */1 * * * *"
}

// Run 任务执行体
func (*RemindRunner) Run() {
	fmt.Println(time.Now())
}
