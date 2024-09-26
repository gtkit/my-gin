package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() {
	config.Add("jwt", func() map[string]any {
		return map[string]any{
			"secret": viper.GetString("jwt.secret"),
		}
	})
}
