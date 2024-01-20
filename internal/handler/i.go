package handler

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/repository"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Ping(c *gin.Context)
	DoPing(c *gin.Context)
}

type handler struct {
	repository repository.Reposit
}

func New(dao dao.Dao) Handler {
	return &handler{
		repository: repository.New(dao.Mdb(), dao.Rdbs()),
	}
}
func (h *handler) i() {}
