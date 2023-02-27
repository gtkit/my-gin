package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(requestid.New())
	// 注册zap相关中间件
	r.Use(GinLogger())       // 日志
	r.Use(GinRecovery(true)) // 错误处理

	// r.Use(Sentinel())
}
