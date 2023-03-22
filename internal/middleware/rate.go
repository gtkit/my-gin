package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gtkit/logger"

	"ydsd_gin/config"

	"github.com/gtkit/golimit"
)

// IP限流器
func LimitIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 得到ip地址
		// ipAddr := ip.GetRealIp(c)
		ipAddr := c.ClientIP()

		if !golimit.Allow(ipAddr, config.GetInt("office.rate")) {
			logger.Warn("ip_warn:", ipAddr, " ip 请求太频繁了, 当前限制每秒: ", config.GetInt("office.rate"))
			// c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 110,
				"msg":  "该IP接口请求太频繁了",
			})
			return
		} else {
			c.Next()
		}
	}
}

// 微信登录状态查询，根据 scene_id 限流
func LimitLoginScene() gin.HandlerFunc {
	return func(c *gin.Context) {
		scene_id := c.Query("scene_id")

		if scene_id == "0" || scene_id == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 110,
				"msg":  "未登录110",
			})
			return
		}

		if !golimit.Allow(scene_id, config.GetInt("office.rate")) {
			logger.Warn("sceneid_warn:", scene_id, " scene_id请求太频繁了")
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 110,
				"msg":  "接口请求太频繁了",
			})
			return
		} else {
			c.Next()
		}
	}
}

// 用户mac地址限流，即 uuid的值
func LimitUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uuid string

		switch method := c.Request.Method; method {
		case "GET":
			uuid = c.Query("uuid")
		case "POST":
			uuid = c.PostForm("uuid")

		}

		if !golimit.Allow(uuid, config.GetInt("office.rate")) {
			logger.Warn("uuid_warn:", uuid, " uuid请求太频繁了")
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 110,
				"msg":  "接口请求太频繁了",
			})
			return
		} else {
			c.Next()
		}
	}
}

// func isAllow(param string, num int) bool {
// 	// return limiter.RateLimiter.GetLimiter(param).Allow()
// 	return limiter.NewLimiter(param, num).Allow()
// }
