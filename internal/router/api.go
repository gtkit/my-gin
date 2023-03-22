package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gtkit/logger"

	"ydsd_gin/internal/apis"
	"ydsd_gin/internal/dao"
)

func ApiRouter(r *gin.RouterGroup) {

	apictl := apis.New(dao.DB(), logger.NewSugar())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apictl.Ping) // ping 接口

	}

}
