package news

import (
	"strings"
	"time"

	"github.com/gtkit/logger"
	fsnews "github.com/gtkit/news"

	"ydsd_gin/config"
	"ydsd_gin/internal/pkg/env"
)

// Warn sends a warning message to the specified news channel.
func Warn(msg string) {
	fsurl := config.GetString("news.fsurl")
	if fsurl == "" {
		return
	}
	t := time.Now().Format(time.DateTime)

	var builder strings.Builder
	builder.WriteString(env.Name())
	builder.WriteString("🕐 时间：")
	builder.WriteString(t)
	builder.WriteString("\n")
	builder.WriteString(msg)
	msg = builder.String()

	fsnews.FsNew(fsurl).Send(msg)
}

// ErrRecord logs an error message and sends it to the specified news channel.
func ErrRecord(msg string) {
	logger.ZError(msg)
	Warn(msg)
}
