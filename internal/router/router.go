// @Author xiaozhaofu 2023/2/17 19:11:00
package router

import (
	"github.com/gin-gonic/gin"

	"ydsd_gin/internal/middleware"
	"ydsd_gin/internal/pkg/env"
	"ydsd_gin/internal/pkg/response"
	v1 "ydsd_gin/internal/router/v1"
)

func InitRouter() *gin.Engine {
	// 产品环境打开此配置, 关闭 debug 模式
	if env.Active().IsPro() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// 不允许的请求方法
	r.HandleMethodNotAllowed = true // 要设置这一步 NoMethod 才生效
	r.NoMethod(nomethod)
	// 未找到的路由
	r.NoRoute(noroute)
	// 产品环境 加载中间件
	middleware.InitMiddleware(r)

	// 注册系统路由
	InitSysRouter(r)

	return r
}

func InitSysRouter(r *gin.Engine) {

	g := r.Group("/draw")
	// 各个路由组
	{
		v1.ApiRouter(g)
	}

}

func noroute(c *gin.Context) {
	response.NotFoundError(c, "未知的路由未知的路由")
}

func nomethod(c *gin.Context) {
	response.NotAllowedMethod(c, "未知的请求方法未知的请求方法")
}
