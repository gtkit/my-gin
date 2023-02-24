package dao

import (
	"sync"

	"gitlab.superjq.com/go-tools/redis/rdb"

	"ydsd_gin/config"
)

// redisClientConfig redis 链接配置信息
type redisClientConfig struct {
	Host     string
	Password string
	prefix   string
	DB       int
}

// RedisConfigs 分组配置信息
type RedisConfigs map[int]*redisClientConfig

// once 确保全局Redis对象只实例一次
var once sync.Once

// redisCollections redis对象集合
var redisCollections map[int]*rdb.Redisclient

func initRedis() map[int]*rdb.Redisclient {
	// 配制 redis 不同的库
	redisConfigs := setRedisConfigs(config.GetIntSlice("redis.dbs"))

	connectRedis(redisConfigs)

	return redisCollections
}

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func connectRedis(configs RedisConfigs) {
	once.Do(func() {
		if redisCollections == nil {
			redisCollections = make(map[int]*rdb.Redisclient, len(redisCollections))
		}

		for dbname, rdbconfig := range configs {
			redisCollections[dbname] = rdb.NewRedis(rdbconfig.Host, rdbconfig.Password, rdbconfig.prefix, rdbconfig.DB)
		}
	})
}

func setRedisConfigs(dbs []int) RedisConfigs {
	redisConfigs := make(RedisConfigs)

	for _, db := range dbs {
		redisConfigs[db] = &redisClientConfig{
			config.GetString("redis.addr"),
			config.GetString("redis.password"),
			config.GetString("redis.db_prefix"),
			db,
		}
	}

	return redisConfigs
}
