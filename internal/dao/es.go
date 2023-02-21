package dao

import (
	"github.com/olivere/elastic/v7"
	"gitlab.superjq.com/go-tools/goes"

	"ydsd_gin/config"
)

func initEsClient() *elastic.Client {
	return goes.New(&goes.Option{
		Host:  config.GetString("es.host"),
		Port:  config.GetString("es.port"),
		User:  config.GetString("es.user"),
		Pass:  config.GetString("es.pass"),
		Debug: config.GetInt("elasticsearch.debug"),
	})

}
