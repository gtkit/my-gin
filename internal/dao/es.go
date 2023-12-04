// @Author xiaozhaofu 2023/7/25 10:50:00
package dao

import (
	"github.com/gtkit/goes"
	"github.com/gtkit/logger"
	"github.com/olivere/elastic/v7"

	"ydsd_gin/config"
)

func initEsClient() *elastic.Client {
	return goes.New(&goes.Option{
		Host:   config.GetString("es.host"),
		Port:   config.GetString("es.port"),
		User:   config.GetString("es.user"),
		Pass:   config.GetString("es.pass"),
		Scheme: config.GetString("es.scheme"),
		Debug:  config.GetInt("elasticsearch.debug"),
		Log:    logger.EsLogger(), // 默认使用 zap log
	})

}
