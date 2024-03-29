package middleware

import (
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gtkit/logger"
	"go.uber.org/zap"
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
		logger.ZInfo("[Gin-Logger]",
			zap.String("startTime", startTime.Format(time.DateTime)),
			zap.String("reqMethod", reqMethod),
			zap.String("requestId", requestid.Get(c)),
			zap.Int("statusCode", statusCode),
			// zap.String("latencyTime", fmt.Sprintf("%v", latencyTime)),
			zap.Duration("latencyTime", latencyTime),
			zap.String("clientIP", clientIP),
			zap.String("reqUri", reqUri),
			zap.String("query", query),
			zap.String("userAgent", userAgent),
		)

		// log.Infof(" %s %s %3d %13v %15s %s %s %s %s",
		// 	startTime.Format("2006-01-02 15:04:05.9999"),
		// 	requestid.Get(c),
		// 	statusCode,
		// 	latencyTime,
		// 	clientIP,
		// 	reqMethod,
		// 	reqUri,
		// 	query,
		// 	userAgent,
		// )
	}

}
