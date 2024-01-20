package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/gtkit/goerr"

	"github.com/gtkit/logger"

	"ydsd_gin/internal/dao"

	"github.com/gtkit/golimit"
)

// IP限流器
func LimitIp(num int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 得到ip地址
		ipAddr := c.ClientIP()
		ipkey := ipAddr + c.Request.URL.Path
		// 限制每秒请求次数
		limit := golimit.NewLimiter(ipkey, num)
		// 判断是否超过限制
		if !limit.Allow() {
			logger.Warn("ip_warn:", ipAddr, " ip 请求太频繁了, 当前限制每秒/ ", num)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": goerr.ErrTooManyRequests.Code,
				"msg":  goerr.ErrTooManyRequests.Desc,
				"data": [0]int{},
			})
			return
		} else {
			c.Next()
		}
	}
}

func RateLimit(num int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 得到ip地址
		ipAddr := c.ClientIP()
		ipkey := ipAddr + c.Request.URL.Path

		if dao.DB().Rdb(1) == nil {
			logger.Warn("limit 未初始化 redis: ", 1)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": goerr.ErrTooManyRequests.Code,
				"msg":  goerr.ErrTooManyRequests.Desc,
				"data": [0]int{},
			})
			return
		}

		limiter := redis_rate.NewLimiter(dao.DB().Rdb(0).Client())
		res, err := limiter.Allow(c.Request.Context(), ipkey, redis_rate.PerSecond(num))
		if err != nil {
			logger.Warn("ip_warn:", ipAddr, " ip 请求太频繁了, 当前限制每秒/ ", num)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": goerr.ErrTooManyRequests.Code,
				"msg":  goerr.ErrTooManyRequests.Desc,
				"data": [0]int{},
			})
			return
		}
		if res.Remaining == 0 {
			logger.Warn("ip_warn:", ipAddr, " ip 请求太频繁了, 当前限制每秒/ ", num)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": goerr.ErrTooManyRequests.Code,
				"msg":  goerr.ErrTooManyRequests.Desc,
				"data": [0]int{},
			})
			return
		}
		c.Next()
	}
}
