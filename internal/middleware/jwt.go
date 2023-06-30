package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gtkit/goerr"

	"github.com/gtkit/encry/jwt"

	"ydsd_gin/config"
	"ydsd_gin/internal/pkg/response"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("App-Token")
		if token == "" {
			response.Error(c, goerr.New(nil, goerr.ErrAuthentication, "请求未携带token，无权限访问"))
			c.Abort()
			return
		}

		jwt.SetSignKey(config.GetString("jwt.secret"))
		j := jwt.NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				response.Error(c, goerr.New(err, goerr.ErrAuthExpired, "授权已过期"))
				c.Abort()
				return
			}
			response.Error(c, goerr.New(err, goerr.ErrAuthentication, "token解析失败"))
			c.Abort()
			return
		}
		// fmt.Println("laravel prv:", claims.Prv)

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Next()

	}
}
