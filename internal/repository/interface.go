// @Author xiaozhaofu 2023/2/23 14:31:00
package repository

import (
	"github.com/gtkit/goerr"
	"github.com/redis/go-redis/v9"

	"my_gin/internal/dao"
	"my_gin/internal/model"
)

func _() {
	var _ Reposit = (*reposit)(nil)
}

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

func (r *reposit) mdb() dao.MyDB {
	return r.dao.Mdb()
}

func (r *reposit) rdb(db int) dao.RDB {
	return r.dao.Rdb(db)
}

func (r *reposit) rdbClient(db int) *redis.Client {
	return r.dao.Rdb(db).Client()
}
