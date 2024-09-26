package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() {
	config.Add("rsa", func() map[string]any {
		return map[string]any{
			"pubpem": viper.GetString("rsa.pub_pem"),
			"pripem": viper.GetString("rsa.pri_pem"),
		}
	})
}
