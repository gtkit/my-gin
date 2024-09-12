package job

import (
	"context"
	"strings"

	"github.com/gtkit/goerr"
	"github.com/hibiken/asynq"
)

const (
	lenTaskType = 2
)

// ReportJobHandler 上报事件.
func ReportJobHandler(ctx context.Context, t *asynq.Task) error {
	reqerr := make(chan error, 1)

	go func() {
		reqerr <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-reqerr:
		return nil
	}
}

// ClickJobHandler 点击事件.
func ClickJobHandler(ctx context.Context, t *asynq.Task) error {
	reqerr := make(chan error, 1)

	go func() {
		reqerr <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-reqerr:
		return nil
	}
}

func taskType(t string) (string, error) {
	tt := strings.Split(t, ":")

	if len(tt) != lenTaskType {
		return "", goerr.Err("task type format error")
	}

	return tt[0], nil
}
