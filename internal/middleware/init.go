package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {

	r.Use(GinLogger())         // 日志
	r.Use(GinRecovery(true))   // 错误处理
	r.Use(TimeoutMiddleware()) // 此中间件要放在 requestid中间件前面, 否则c.write中的requestid会被timeout中间件清空
	r.Use(requestid.New())
	// r.Use(Sentinel())
}
