package news

import (
	"time"

	"github.com/gtkit/logger"
	gtnews "github.com/gtkit/news"
	"github.com/gtkit/stringx"

	"my_gin/config"
	"my_gin/internal/pkg/env"
)

// Warn sends a warning message to the specified news channel.
func Warn(msg string, url ...string) {
	var fsurl string
	if len(url) > 0 {
		fsurl = url[0]
	} else {
		fsurl = config.GetString("news.fsurl")
	}

	t := time.Now().Format(time.DateTime)

	info := stringx.BuilderJoin([]string{
		env.Name(),
		"🕐 时间：",
		t,
	})

	gtnews.FsNew(fsurl).Send(info, msg)
}

// ErrRecord logs an error message and sends it to the specified news channel.
func ErrRecord(msg string) {
	logger.ZError(msg)
	Warn(msg)
}
