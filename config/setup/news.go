package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func addnews() {
	config.Add("news", func() map[string]interface{} {
		return map[string]interface{}{
			"fsurl": viper.GetString("news.fsurl"),
		}
	})

}
