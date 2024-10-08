package setup

import (
	"github.com/spf13/viper"

	"my_gin/config"
)

func init() {
	config.Add("news", func() map[string]any {
		return map[string]any{
			"fsurl": viper.GetString("news.fsurl"),
		}
	})
}
