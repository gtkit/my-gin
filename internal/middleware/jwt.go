package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/gtkit/encry/jwt"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("App-Token")
		if token == "" {
			c.JSON(401, gin.H{
				"errcode": 22006,
				"errmsg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// logger.Infof("get token: ", token)

		j := jwt.NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(401, gin.H{
					"errdoce": 22007,
					"errmsg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(401, gin.H{
				"errcode": 22008,
				"errmsg":  err.Error(),
			})
			c.Abort()
			return
		}
		// fmt.Println("laravel prv:", claims.Prv)

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)

	}
}
