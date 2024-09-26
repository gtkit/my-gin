package task

import (
	"sync"
	"time"

	"github.com/gtkit/logger"
	"github.com/hibiken/asynq"

	"my_gin/config"
	"my_gin/tools/utils"
)

const (
	MaxRetryCount  = 3
	retryDelay     = 5 * time.Second
	TimeOut        = 10 * time.Second
	ConcurrencyNum = 100 // 并发数
	Priority       = 1   // 优先级
)

var (
	client     *asynq.Client
	server     *asynq.Server
	clientOnce sync.Once
	serverOnce sync.Once
)

type Option []asynq.Option

type Info *asynq.TaskInfo

func Server() *asynq.Server {
	serverOnce.Do(func() {
		server = asynq.NewServer(
			redisOpt(),
			asynqConfig(),
		)
	})
	return server
}

func NewClient() {
	clientOnce.Do(func() {
		client = asynq.NewClient(redisOpt())
	})
}

func Client() *asynq.Client {
	if client == nil {
		NewClient()
	}
	return client
}

func Close() error {
	return Client().Close()
}

func NewServeMux() *asynq.ServeMux {
	return asynq.NewServeMux()
}

func StopServe(srv *asynq.Server) {
	srv.Stop()
	srv.Shutdown()
}

func MaxRetry(n int) asynq.Option {
	return asynq.MaxRetry(n)
}

func NewTask(typename string, payload []byte, opts ...asynq.Option) *asynq.Task {
	return asynq.NewTask(typename, payload, opts...)
}

func ID(id string) asynq.Option {
	return asynq.TaskID(id)
}

func Queue(queue string) asynq.Option {
	return asynq.Queue(queue)
}

func Timeout(t time.Duration) asynq.Option {
	return asynq.Timeout(t)
}

func ProcessIn(d time.Duration) asynq.Option {
	return asynq.ProcessIn(d)
}

func UniqueID(mediaType string) string {
	return mediaType + ":" + utils.UUID()
}

func TaskOption(mediaType string) Option {
	return Option{
		MaxRetry(MaxRetryCount),
		Queue(mediaType + ":queue"),
		Timeout(TimeOut),
	}
}

func redisOpt() asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     config.GetString("redis.addr"),
		Password: config.GetString("redis.password"),
		DB:       config.GetIntSlice("redis.dbs")[0],
	}
}

func asynqConfig() asynq.Config {
	return asynq.Config{
		Concurrency: ConcurrencyNum,
		Queues:      QueueLists(),
		Logger:      logger.Sugar(),
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return retryDelay
		},
	}
}

// QueueLists 队列列表.
func QueueLists() map[string]int {
	queues := make(map[string]int)

	return queues
}
