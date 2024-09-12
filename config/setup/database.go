package setup

import (
	"ydsd_gin/config"

	"github.com/spf13/viper"
)

func init() {
	config.Add("database", func() map[string]any {
		return map[string]any{
			"dbType":         viper.GetString("database.dbtype"),
			"host":           viper.GetString("database.host"),
			"port":           viper.GetString("database.port"),
			"name":           viper.GetString("database.name"),
			"username":       viper.GetString("database.username"),
			"password":       viper.GetString("database.password"),
			"maxOpenConn":    viper.GetInt("database.maxOpenConn"),
			"maxIdleConn":    viper.GetInt("database.maxIdleConn"),
			"maxlifeseconds": viper.GetInt("database.maxlifeseconds"),
		}
	})
}
