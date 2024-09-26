// @Author xiaozhaofu 2023/3/23 14:58:00
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"

	"github.com/gin-contrib/timeout"

	resp "my_gin/internal/pkg/response"
)

const (
	DefaultTimeout = 30 * time.Second
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(DefaultTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			// 在此恢复请求头
			// c.Writer.Header().Set("X-Request-Id", c.GetHeader("X-Request-Id"))
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)

}

func timeoutResponse(c *gin.Context) {
	resp.Error(c, goerr.New(nil, goerr.Timeout(), "http 请求超时"))
}
