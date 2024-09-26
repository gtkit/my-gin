package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() {
	config.Add("app", func() map[string]any {
		return map[string]any{
			"readTimeout":   viper.GetInt("application.readTimeout"),
			"writerTimeout": viper.GetInt("application.writerTimeout"),
			"host":          viper.GetString("application.host"),
			"port":          viper.GetString("application.port"),
			"name":          viper.GetString("application.name"),
			"mode":          viper.GetString("application.mode"),
			"domain":        viper.GetString("application.domain"),
			"ishttps":       viper.GetBool("application.ishttps"),
		}
	})
}
