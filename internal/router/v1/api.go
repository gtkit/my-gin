package v1

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/handler"
)

func ApiRouter(r *gin.RouterGroup) {

	apictl := handler.New(dao.DB())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apictl.Ping)    // ping 接口
		v1.POST("/ping", apictl.DoPing) // ping 接口

	}

}
