package cron

import (
	"github.com/gtkit/logger"
	"github.com/robfig/cron/v3"
)

func New() *cron.Cron {
	return cron.New(cron.WithSeconds(), cron.WithLogger(logger.CronLog()))
}

// 停止任务.
func StopCron(c *cron.Cron, el cron.EntryID) {
	c.Remove(el)
	c.Stop()
}

type Job struct {
}

func NewJob() *Job {
	return &Job{}
}
func (j *Job) Run() {

}
