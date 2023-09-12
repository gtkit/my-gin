package v1

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/apis"
	"ydsd_gin/internal/dao"
)

func ApiRouter(r *gin.RouterGroup) {

	apictl := apis.New(dao.DB())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apictl.Ping) // ping 接口

	}

}
