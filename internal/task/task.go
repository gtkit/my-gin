// @Author xiaozhaofu 2023/5/23 20:53:00
package task

import (
	"context"

	"github.com/gtkit/asynqx"
	"github.com/gtkit/logger"

	"my_gin/internal/pkg/asynq"
)

type Job struct {
}

func New() *Job {
	return &Job{}
}
func (j *Job) Run() {
	// 这里写要运行的业务逻辑

}

func Srv() *asynqx.Server {
	return asynq.Srv()
}

func HandleTaskWithCtx(ctx context.Context, taskType string, taskData *any) error {
	if taskData == nil {
		logger.Info("Task data is nil")
		return nil
	}
	return nil
}
