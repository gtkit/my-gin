package dao

import (
	"gitlab.superjq.com/go-tools/redis/rdb"

	"ydsd_gin/config"
)

func initRedis() *rdb.Redisclient {
	// 初始化链接 确保全局的 Redis 对象只实例一次

	// 赋值给全局变量
	return rdb.NewRedis(
		config.GetString("redis.addr"),
		config.GetString("redis.password"),
		config.GetString("redis.db_prefix"),
		config.GetInt("redis.db"),
	)

}
