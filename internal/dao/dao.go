// @Author xiaozhaofu 2023/2/20 10:52:00
package dao

import (
	"github.com/gtkit/redis/rdb"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"ydsd_gin/config"
)

var daoDB Dao

var _ Dao = (*dao)(nil)

type Dao interface {
	Mdb() *gorm.DB                  // mysql 数据库
	Rdb(db int) *rdb.Redisclient    // 获取指定库的redis
	Rdbs() map[int]*rdb.Redisclient // 获取所有redis连接
	MdbClose() error                // 关闭 mysql 连接
	RdbClose() error                // 关闭 redis 连接
	ES() *elastic.Client            // elasticsearch 实例
	d()                             // 防止被其他包实现
}
type dao struct {
	rdb map[int]*rdb.Redisclient // redis
	mdb *gorm.DB                 // gorm mysql
	es  *elastic.Client
}

func (d *dao) Mdb() *gorm.DB {
	return d.mdb
}

func (d *dao) Rdb(db int) *rdb.Redisclient {
	return d.rdb[db]
}

func (d *dao) Rdbs() map[int]*rdb.Redisclient {
	return d.rdb
}

func (d *dao) MdbClose() error {
	db, err := d.mdb.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (d *dao) RdbClose() error {
	for _, v := range config.GetIntSlice("redis.dbs") {
		if d.rdb[v] == nil {
			continue
		}
		err := d.rdb[v].Client().Close()

		if err != nil {
			return err
		}
	}
	return nil
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
