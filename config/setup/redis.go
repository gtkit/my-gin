package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func addredis() {
	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{
			"addr":      viper.GetString("redis.addr"),
			"password":  viper.GetString("redis.password"),
			"dbs":       viper.GetIntSlice("redis.dbs"),
			"db_prefix": viper.GetString("redis.db_prefix"),
		}
	})
}
