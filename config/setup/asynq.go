package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() { //nolint:gochecknoinits // this is why
	config.Add("asynq", func() map[string]any {
		return map[string]any{
			"redisAddr":     viper.GetString("asynq.redis.addr"),
			"redisUsername": viper.GetString("asynq.redis.username"),
			"redisPassword": viper.GetString("asynq.redis.password"),
			"redisDB":       viper.GetInt("asynq.redis.db"),
			"concurrency":   viper.GetInt("asynq.config.concurrency"),
			"maxRetries":    viper.GetInt("asynq.config.maxretries"),
		}
	})
}
