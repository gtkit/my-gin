// @Author xiaozhaofu 2023/3/23 14:58:00
package middleware

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"

	"ydsd_gin/internal/pkg/response"
)

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(30*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)

}

func timeoutResponse(c *gin.Context) {
	response.Error(c, goerr.ErrTimeout, goerr.Err("http 请求超时"))
}
