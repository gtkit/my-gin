package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() {
	config.Add("redis", func() map[string]any {
		return map[string]any{
			"addr":     viper.GetString("redis.addr"),
			"password": viper.GetString("redis.password"),
			"dbs":      viper.GetIntSlice("redis.dbs"),
			"prefix":   viper.GetString("redis.prefix"),
		}
	})
}
