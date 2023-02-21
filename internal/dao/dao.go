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
	Es() *elastic.Client
	Close() error
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

func (d *dao) Es() *elastic.Client {
	return d.es
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
		es:  initEsClient(),
	}

}

func DB() Dao {
	return daoDB
}

func Mdb() *gorm.DB {
	return DB().Mdb()
}

func Rdb() *rdb.Redisclient {
	return DB().Rdb()
}
