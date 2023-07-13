package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gtkit/logger"

	"ydsd_gin/config"
	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/router"
)

func Run() {
	// 初始化路由
	r := router.InitRouter()

	defer close()

	srv := &http.Server{
		Addr:    config.GetString("app.host") + ":" + config.GetString("app.port"),
		Handler: r,
	}

	// 启动服务
	go startServe(srv)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1) // 定义管道来装信号
	// 设置需要接受哪些信号量才能通过管道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format("2006/01/02 15:04:05"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")

}

func startServe(srv *http.Server) {
	// 服务连接
	if config.GetBool("app.ishttps") {
		logger.Infof("\u001B[32m%s\u001B[0m", "https 服务启动----->>> "+config.GetString("app.host")+":"+config.GetString("app.port"))
		if err := srv.ListenAndServeTLS(config.GetString("ssl.pem"), config.GetString("ssl.key")); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
			return
		}

	} else {
		logger.Infof("\u001B[32m%s\u001B[0m", "http 服务启动----->>> "+config.GetString("app.host")+":"+config.GetString("app.port"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
			return
		}

	}
}

func close() {
	dbclose()
	redisclose()
}
func dbclose() {
	err := dao.DB().MdbClose()
	if err != nil {
		logger.Error("mysql close  error: ", err)
		return
	}
	logger.Info(" mysql close  success")

}
func redisclose() {
	err := dao.DB().RdbClose()
	if err != nil {
		logger.Error("redis close  error: ", err)
		return
	}
	logger.Info(" redis close success ")
}
