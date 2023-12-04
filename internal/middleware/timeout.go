// @Author xiaozhaofu 2023/3/23 14:58:00
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"

	"github.com/gin-contrib/timeout"

	resp "ydsd_gin/internal/pkg/response"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(30*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			// 在此恢复请求头
			// c.Writer.Header().Set("X-Request-Id", c.GetHeader("X-Request-Id"))
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)

}

func timeoutResponse(c *gin.Context) {
	resp.Error(c, goerr.New(nil, goerr.ErrTimeout, "http 请求超时"))
}
