package router

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/apis"
	"ydsd_gin/internal/pkg/response"
)

func ApiRouter(r *gin.RouterGroup) {

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", response.Decorate(apis.Ping)) // ping 接口

	}

}
