// @Author xiaozhaofu 2023/2/20 10:52:00
package dao

import (
	"gitlab.superjq.com/go-tools/redis/rdb"
	"gorm.io/gorm"
)

var daoDB Dao

var _ Dao = (*dao)(nil)

type Dao interface {
	Mdb() *gorm.DB         // mysql 数据库
	Rdb() *rdb.Redisclient // redis
	Close() error
	d()
}
type dao struct {
	rdb *rdb.Redisclient // redis
	mdb *gorm.DB         // gorm mysql

}

func (d *dao) Mdb() *gorm.DB {
	return d.mdb
}

func (d *dao) Rdb() *rdb.Redisclient {
	return d.rdb
}
func (d *dao) Close() error {
	db, err := d.mdb.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	err = d.rdb.Client().Close()
	if err != nil {
		return err
	}
	return nil
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
