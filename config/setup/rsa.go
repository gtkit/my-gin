package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func init() {
	config.Add("rsa", func() map[string]interface{} {
		return map[string]interface{}{
			"pubpem": viper.GetString("rsa.pub_pem"),
			"pripem": viper.GetString("rsa.pri_pem"),
		}
	})
}
