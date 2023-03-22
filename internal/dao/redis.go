package dao

import (
	"github.com/gtkit/redis/rdb"

	"ydsd_gin/config"
)

func initRedisCollection() map[int]*rdb.Redisclient {

	return rdb.NewRedisCollection(
		config.GetString("redis.addr"),     // redis 地址
		config.GetString("redis.password"), // redis密码
		config.GetString("redis.prefix"),   // redis键值前缀
		config.GetIntSlice("redis.dbs"),    // 用到的redis 多个库, 一个库时,只需设置需要用到库
	)
}

// func initRedis() *rdb.Redisclient {
// 	// 赋值给全局变量
// 	return rdb.NewRedis(
// 		config.GetString("redis.addr"),
// 		config.GetString("redis.password"),
// 		config.GetString("redis.prefix"),
// 		config.GetInt("redis.db"), // 指定单个库
// 	)
//
// }
