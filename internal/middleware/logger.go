package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.superjq.com/go-tools/logger"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求参数
		query := c.Request.URL.RawQuery

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 请求User-Agent
		userAgent := c.GetHeader("User-Agent")

		// 日志格式
		logger.Infof(" %s %3d %13v %15s %s %s %s %s",
			startTime.Format("2006-01-02 15:04:05.9999"),
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			query,
			userAgent,
		)
	}

}