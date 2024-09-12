// @Author xiaozhaofu 2023/7/25 11:17:00
package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func init() {
	config.Add("es", func() map[string]any {
		return map[string]any{
			"host":   viper.GetString("elasticsearch.host"),
			"port":   viper.GetString("elasticsearch.port"),
			"user":   viper.GetString("elasticsearch.user"),
			"pass":   viper.GetString("elasticsearch.pass"),
			"scheme": viper.GetString("elasticsearch.scheme"),
			"debug":  viper.GetInt("elasticsearch.debug"),
		}
	})
}
