package logger

import (
	"gitlab.superjq.com/go-tools/logger"

	"ydsd_gin/config"
)

func Init() {
	opt := &logger.Option{
		Level:         config.GetString("log.level"),
		ConsoleStdout: config.GetBool("log.consolestdout"),
		FileStdout:    config.GetBool("log.filestdout"),
		Division:      config.GetString("log.division"),
		Path:          config.GetString("log.path"),
		SqlLog:        config.GetBool("log.sql"),
	}
	logger.NewZap(opt)

}
