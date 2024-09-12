// @Author xiaozhaofu 2023/2/23 14:31:00
package repository

import (
	"github.com/gtkit/goerr"
	"github.com/gtkit/redis"
	"gorm.io/gorm"

	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/model"
)

var _ Reposit = (*reposit)(nil)

type Reposit interface {
	i()
	Ping() (*model.AssistantMember, goerr.Error)
}
type reposit struct {
	dao dao.Dao
}

func New(dao dao.Dao) Reposit {
	return &reposit{
		dao: dao,
	}
}

func (r *reposit) i() {}

func (r *reposit) rdb(db int) *redis.Redisclient {
	return r.dao.Rdb(db)
}

func (r *reposit) mydb() *gorm.DB {
	return r.dao.Mdb()
}
