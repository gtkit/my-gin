package api

import (
	"github.com/gin-gonic/gin"

	"my_gin/internal/dao"
	"my_gin/internal/repository"
)

func _() {
	var _ Handler = (*handler)(nil) // 强制 *handler 去实现 Handler 接口，编译器会检查 *handler 类型是否实现了 Handler 接口。
}

type Handler interface {
	i() // 强制 Handler 接口中所有方法只能在本包中去实现，在其他包中不允许去实现。因为接口中有小写方法，所以在其他包无法去实现。i() 表示一个小写方法，起其他名字也可以
	Ping(c *gin.Context)
}

type handler struct {
	repository repository.Reposit
}

func New(dao dao.Dao) Handler {
	return &handler{
		repository: repository.New(dao),
	}
}
func (h *handler) i() {
	// TODO implement me
	panic("implement me")
}
