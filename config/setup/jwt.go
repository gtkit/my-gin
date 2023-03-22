package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			"  secret": viper.GetString("jwt.secret"),
		}
	})
}
