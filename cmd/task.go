package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gtkit/logger"

	"my_gin/task/cron"

	"github.com/spf13/cobra"

	"my_gin/config"
	"my_gin/internal/task"
)

// taskCmd represents the task command.
var taskCmd = &cobra.Command{
	Use: "task",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Greenf("asynq task enqueue start ...", time.Now().Format(time.DateTime))
		ctx, cancel := context.WithCancel(context.Background())
		// 任务协程
		go dotask(ctx)

		quit := make(chan os.Signal, 1) // 定义管道来装信号
		// 设置需要接受哪些信号量才能通过管道
		signal.Notify(quit, osSignal()...)
		<-quit
		cancel()
		logger.Redf("asynq task enqueue done ...", time.Now().Format(time.DateTime))
	},
}

func init() { //nolint:gochecknoinits // this is why // this is why
	rootCmd.AddCommand(taskCmd)
}

// 创建任务队列.
func dotask(ctx context.Context) {
	// 可以在这里启动定时任务,自定义日志
	c := cron.New()
	c.Start()

	el, err := c.AddJob(config.GetString("app.cron"), task.New())
	if err != nil {
		logger.Error("AddJob error: ", err.Error())
		panic(err)
	}
	defer cron.StopCron(c, el)
	<-ctx.Done()
	logger.Info("cron task job stop")
}

// osSignal returns a list of signals to listen for.
func osSignal() []os.Signal {
	return []os.Signal{
		// syscall.SIGTSTP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
}
