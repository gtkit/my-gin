// @Author xiaozhaofu 2023/5/14 13:56:00
package middleware

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gtkit/encry/jwt"
	"github.com/gtkit/goerr"
	"github.com/gtkit/logger"

	"my_gin/config"
	"my_gin/internal/model"
)

const AuthErrCode = 401

// SseHeader sse header
func SseHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

// SseAuth sse auth
func SseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var err error

		// sse 连接方式不支持自定义header，只能通过url传递token或者post参数传递token
		token = c.Query("app_token")
		if token == "" {
			token = c.PostForm("app_token")
			if token == "" {
				info := model.SseMsg{
					Code: AuthErrCode,
					Msg:  err.Error(),
				}
				SentEvent(c, "error", info)
				logger.Error("-------app_token: ", err.Error())
				c.Abort()
				return
			}
		}

		if token == "" {
			// response.Error(c, goerr.ErrAuthentication, goerr.Err("请求未携带token，无权限访问"))
			info := model.SseMsg{
				Code: AuthErrCode,
				Msg:  "请求未携带token，无权限访问",
			}
			SentEvent(c, "error", info)
			c.Abort()
			return
		}
		jwt.SetSignKey(config.GetString("jwt.secret"))
		j := jwt.NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if goerr.Is(err, jwt.ErrTokenExpired) {
				info := model.SseMsg{
					Code: AuthErrCode,
					Msg:  "授权已过期, 请重新登录",
				}
				SentEvent(c, "error", info)
				c.Abort()
				return
			}
			data := model.SseMsg{
				Code: 401,
				Msg:  "授权已过期, 请重新登录 2",
			}

			SentEvent(c, "error", data)

			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		// c.Set("claims", claims)
		c.Set("userid", claims.Subject)
		c.Next()
	}
}

func SentEvent(c *gin.Context, event string, data model.SseMsg) {
	c.Stream(func(w io.Writer) bool {
		c.SSEvent(event, data)
		return false
	})
}
