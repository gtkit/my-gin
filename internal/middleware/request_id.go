package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"ydsd_gin/tools/utils"
)

const (
	TrafficKey = "X-Request-Id"
	LoggerKey  = "_go-admin-logger-request"
)

// RequestId 自动增加requestId
func RequestId(trafficKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}
		requestId := c.GetHeader(trafficKey)
		if requestId == "" {
			requestId = c.GetHeader(strings.ToLower(trafficKey))
		}
		if requestId == "" {
			requestId = utils.NewUuid()
		}
		c.Request.Header.Set(trafficKey, requestId)
		c.Set(trafficKey, requestId)

		c.Next()
	}
}
