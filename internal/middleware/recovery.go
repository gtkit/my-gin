// @Author xiaozhaofu 2023/1/30 15:00:00
package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/logger"
	"go.uber.org/zap"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error:", err),
						zap.String("request:", string(httpRequest)),
					)
					c.JSON(http.StatusOK, gin.H{
						"code": 102,
						"msg":  err.(error).Error(),
					})
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request--", string(httpRequest)),
						zap.String("stack--", string(debug.Stack())),
					)
					c.JSON(http.StatusOK, gin.H{
						"code": 103,
						"msg":  err.(error).Error(),
					})
					c.AbortWithStatus(http.StatusBadRequest)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.JSON(http.StatusOK, gin.H{
						"code": 104,
						"msg":  err.(error).Error(),
					})
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
