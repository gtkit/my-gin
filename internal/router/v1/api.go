package v1

import (
	"github.com/gin-gonic/gin"

	"my_gin/internal/api"
	"my_gin/internal/dao"
)

func ApiRouter(r *gin.RouterGroup) {

	apictl := api.New(dao.DB())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apictl.Ping) // ping 接口
		// v1.Any("/ping", apictl.DoPing)  // ping 接口, any 前不能有相同路由的方法

	}

}
