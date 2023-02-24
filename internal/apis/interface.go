// @Author xiaozhaofu 2023/2/23 11:23:00
package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/repository"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Ping(c *gin.Context)
}

type handler struct {
	log        *zap.SugaredLogger
	repository repository.Reposit
}

func New(dao dao.Dao, log *zap.SugaredLogger) Handler {
	fmt.Println("-----begin new ctrl------")
	return &handler{
		log:        log,
		repository: repository.New(dao.Mdb(), dao.Rdb(1), dao.Rdbs()),
	}
}
func (h *handler) i() {}
