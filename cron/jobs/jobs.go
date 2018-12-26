package jobs

import (
	"github.com/pibigstar/go-todo/utils/logger"
)

var log = logger.New("jobs")

type job interface {
	Run()
	Cron() string
	Name() string
}

var jobs []job

func addJobs(runner job) {
	jobs = append(jobs, runner)
}

// GetJobs 得到所有定时任务
func GetJobs() []job {
	return jobs
}
