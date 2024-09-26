package asynq

import (
	"time"

	"github.com/gtkit/asynqx"

	"github.com/gtkit/logger"

	"my_gin/config"
)

var server *asynqx.Server

func New() {
	server = asynqx.NewServer(
		asynqx.WithRedisAddr(config.GetString("asynq.redisAddr")),
		asynqx.WithRedisDB(config.GetInt("asynq.redisDB")),
		asynqx.WithConcurrency(config.GetInt("asynq.concurrency")),
		asynqx.WithTaskTimeout(10*time.Minute),
		asynqx.WithLogger(logger.Sugar()),
	)
}

func Srv() *asynqx.Server {
	return server
}
