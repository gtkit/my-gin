package log

import (
	"github.com/gtkit/logger"
	"github.com/olivere/elastic/v7"
)

var _ elastic.Logger = (*EsLogger)(nil)

type EsLogger struct {
}

func (l EsLogger) Printf(format string, v ...interface{}) {
	logger.Infof("[--ES] "+format, v...)
}
