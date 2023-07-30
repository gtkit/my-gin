package log

import (
	"github.com/go-resty/resty/v2"
	"github.com/gtkit/logger"
)

// restyLogger resty log
var _ resty.Logger = (*RestyLogger)(nil)

type RestyLogger struct {
}

func (l *RestyLogger) Errorf(format string, v ...interface{}) {
	logger.Errorf("--ERROR RESTY "+format, v)
}
func (l *RestyLogger) Warnf(format string, v ...interface{}) {
	logger.Warnf("--WARN RESTY "+format, v)
}
func (l *RestyLogger) Debugf(format string, v ...interface{}) {
	logger.Debugf("--DEBUG RESTY "+format, v)
}
