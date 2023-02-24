// @Author xiaozhaofu 2023/2/20 10:52:00
package dao

import (
	"github.com/olivere/elastic/v7"
	"gitlab.superjq.com/go-tools/redis/rdb"
	"gorm.io/gorm"
)

var daoDB Dao

var _ Dao = (*dao)(nil)

type Dao interface {
	Mdb() *gorm.DB         // mysql 数据库
	Rdb() *rdb.Redisclient // redis
	MdbClose() error       // 关闭 mysql 连接
	RdbClose() error       // 关闭 redis 连接
	d()
}
type dao struct {
	rdb *rdb.Redisclient // redis
	mdb *gorm.DB         // gorm mysql
	es  *elastic.Client
}

func (d *dao) Mdb() *gorm.DB {
	return d.mdb
}

func (d *dao) Rdb() *rdb.Redisclient {
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
	return d.rdb.Client().Close()
}

func (d *dao) d() {}

func New() {
	daoDB = &dao{
		rdb: initRedis(),
		mdb: initMysql(),
	}
}

func DB() Dao {
	return daoDB
}
