package middleware

import (
	"github.com/alibaba/sentinel-golang/core/system"
	sentinel "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
	"github.com/gtkit/logger"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 20,
			Strategy:     system.BBR,
		},
	}); err != nil {
		logger.Fatalf("Unexpected error: %+v", err)
	}
	return sentinel.SentinelMiddleware(
		sentinel.WithResourceExtractor(func(ctx *gin.Context) string {
			return ctx.GetHeader("X-Real-IP")
		}),
		sentinel.WithBlockFallback(func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(200, map[string]interface{}{
				"msg":  "too many request; the quota used up!",
				"code": 500,
			})
		}),
	)
}
