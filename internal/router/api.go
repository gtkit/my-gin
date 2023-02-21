package router

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/apis"
)

func ApiRouter(r *gin.RouterGroup) {

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", json(apis.Ping)) // ping 接口

	}

}
