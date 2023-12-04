package log

import (
	"github.com/gtkit/logger"

	"ydsd_gin/config"
)

// Init
/**
1. ElasticSearch log 实例: logger.Es().
2. Cron log 实例: logger.Cron()
3. 颜色输出: logger.Yellow(msg), logger.Blue(msg), logger.Green(msg), logger.Red(msg), logger. Redf(msg, arg...)
*/
func Init() {
	logger.NewZap(&logger.Option{
		Level:         config.GetString("log.level"),
		ConsoleStdout: config.GetBool("log.consolestdout"),
		FileStdout:    config.GetBool("log.filestdout"),
		Division:      config.GetString("log.division"),
		Path:          config.GetString("log.path"),
		SqlLog:        config.GetBool("log.sql"),
	})
}
