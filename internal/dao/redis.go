package dao

import (
	"github.com/gtkit/redis"

	"my_gin/config"
)

func initRedisCollection() map[int]*redis.Redisclient {
	var dbsconn []redis.ConnConfigOption
	dbsconn = append(dbsconn, redis.WithAddr(config.GetString("redis.addr")))
	for _, db := range config.GetIntSlice("redis.dbs") {
		dbsconn = append(dbsconn, redis.WithDB(db))
	}

	return redis.NewCollection(
		dbsconn...,
	)
}
