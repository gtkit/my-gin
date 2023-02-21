// @Author xiaozhaofu 2023/2/21 19:43:00
package router

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/pkg/response"
)

type Handler func(ctx *gin.Context) *response.Response

func json(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := h(c)
		if r != nil {
			c.SecureJSON(r.HttpStatus, &r.Res)
		}
		response.PutResponse(r)
	}
}
