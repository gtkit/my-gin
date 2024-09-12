// @Author xiaozhaofu 2023/2/20 10:52:00
package dao

import (
	"github.com/gtkit/redis"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"ydsd_gin/config"
)

var (
	daoDB Dao
	_     Dao = (*dao)(nil)
)

type (
	MyDB = *gorm.DB
	RDB  = *redis.Redisclient
)

type Dao interface {
	Mdb() MyDB           // mysql 数据库
	Rdb(db int) RDB      // 获取指定库的redis
	MdbClose() error     // 关闭 mysql 连接
	RdbClose() error     // 关闭 redis 连接
	ES() *elastic.Client // elasticsearch 实例
	d()                  // 防止被其他包实现
}
type dao struct {
	rdb map[int]RDB // redis,, map[db]client, db从配置文件中读取
	mdb MyDB        // gorm mysql
	es  *elastic.Client
}

func (d *dao) Mdb() MyDB {
	return d.mdb
}

func MdbClose() error {
	return daoDB.MdbClose()
}

func (d *dao) MdbClose() error {
	db, err := d.mdb.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func Rdb(db int) RDB {
	return DB().Rdb(db)
}

func (d *dao) Rdb(db int) RDB {
	return d.rdb[db]
}

func (d *dao) RdbClose() error {
	for _, v := range config.GetIntSlice("redis.dbs") {
		if d.rdb[v] == nil {
			continue
		}
		if err := d.rdb[v].Client().Close(); err != nil {
			return err
		}
	}
	return nil
}
func RdbClose() error {
	return daoDB.RdbClose()
}

func (d *dao) d() {}

func New() {
	daoDB = &dao{
		rdb: initRedisCollection(),
		mdb: initMysql(),
		es:  initEsClient(),
	}
}

func DB() Dao {
	return daoDB
}

// ES elasticsearch 实例
func (d *dao) ES() *elastic.Client {
	return d.es
}
