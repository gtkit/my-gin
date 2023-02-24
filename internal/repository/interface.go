// @Author xiaozhaofu 2023/2/23 14:31:00
package repository

import (
	"gitlab.superjq.com/go-tools/redis/rdb"
	"gorm.io/gorm"

	"ydsd_gin/internal/model"
)

var _ Reposit = (*reposit)(nil)

type Reposit interface {
	i()
	Ping() (*model.AssistantMember, error)
}
type reposit struct {
	mdb  *gorm.DB
	rdb  *rdb.Redisclient
	rdbs map[int]*rdb.Redisclient
}

func New(sqldb *gorm.DB, redis *rdb.Redisclient, rdbs map[int]*rdb.Redisclient) Reposit {
	return &reposit{
		mdb:  sqldb,
		rdb:  redis,
		rdbs: rdbs,
	}
}

func (r *reposit) i() {}
