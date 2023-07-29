package cmd

import (
	"github.com/gtkit/logger"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"

	"ydsd_gin/cmd/server"
	"ydsd_gin/internal/task"
)

const _UI = `
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚═╝     ╚═╝
`

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("\x1b[32m%s\x1b[0m", _UI)
		go dotask()
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func dotask() {
	// 可以在这里启动定时任务,自定义日志
	c := cron.New(cron.WithSeconds(), cron.WithLogger(&tasklog{}))
	c.Start()

	el, err := c.AddJob("@every 120s", task.New())
	if err != nil {
		logger.Error("AddJob error", el, err)
		c.Stop()
	}
	select {}
}

var _ cron.Logger = &tasklog{}

type tasklog struct {
}

func (t *tasklog) Info(msg string, keysAndValues ...interface{}) {
	logger.Infof("[定时任务 INFO]: "+msg+"--", keysAndValues...)
}
func (t *tasklog) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Errorf("[定时任务 ERROR]: "+msg+"--"+err.Error(), keysAndValues...)
}
