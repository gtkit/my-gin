package log

import (
	"github.com/gtkit/logger"
	"github.com/robfig/cron/v3"
)

var _ cron.Logger = &Tasklog{}

type Tasklog struct {
}

func (t *Tasklog) Info(msg string, keysAndValues ...interface{}) {
	logger.Infof("[定时任务 INFO]: "+msg+"--", keysAndValues...)
}
func (t *Tasklog) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Errorf("[定时任务 ERROR]: "+msg+"--"+err.Error(), keysAndValues...)
}
