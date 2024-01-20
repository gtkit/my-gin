package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"

	"github.com/gtkit/encry/jwt"

	"ydsd_gin/internal/pkg/response"
)

// JWTAuth 中间件，检查token

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("App-Token")
		if token == "" {
			resp.Error(c, goerr.New(nil, goerr.ErrAuthentication, "请求未携带token，无权限访问"))
			c.Abort()
			return
		}

		j := jwt.NewJWT() // 唯一的jwt实例
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if goerr.Is(err, jwt.TokenExpired) {
				resp.Error(c, goerr.New(err, goerr.ErrAuthExpired, "授权已过期"))
				c.Abort()
				return
			}

			resp.Error(c, goerr.New(err, goerr.ErrAuthentication, "token解析失败"))
			c.Abort()
			return
		}

		if claims.JwtRole() != "client" {
			resp.Error(c, goerr.New(err, goerr.ErrAuthentication, "角色不对"))
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		// c.Set("claims", claims)
		c.Set("userid", claims.JwtSubject())
		c.Next()
	}
}
