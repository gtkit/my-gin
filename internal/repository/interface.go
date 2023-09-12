// @Author xiaozhaofu 2023/2/23 14:31:00
package repository

import (
	"github.com/gtkit/goerr"
	"github.com/gtkit/redis/rdb"
	"gorm.io/gorm"

	"ydsd_gin/internal/model"
)

var _ Reposit = (*reposit)(nil)

type Reposit interface {
	i()
	Ping() (*model.AssistantMember, goerr.Error)
}
type reposit struct {
	mdb  *gorm.DB
	rdbs map[int]*rdb.Redisclient
}

func New(sqldb *gorm.DB, rdbs map[int]*rdb.Redisclient) Reposit {
	return &reposit{
		mdb:  sqldb,
		rdbs: rdbs,
	}
}

func (r *reposit) i() {}
