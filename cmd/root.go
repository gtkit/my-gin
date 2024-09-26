package cmd

import (
	"fmt"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/gtkit/logger"
	"github.com/gtkit/verify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"my_gin/config"
	"my_gin/config/setup"
	"my_gin/internal/dao"
	"my_gin/internal/pkg/asynq"

	"my_gin/internal/pkg/env"
	jwt "my_gin/internal/pkg/jwtauth"
	"my_gin/internal/pkg/log"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 配置初始化
		setup.Initialize()
		// 初始化 log
		log.Init()
		// 初始化 jwt
		jwt.InitJwt()
		// 链接数据库, redis
		dao.New()
		// 初始化验证翻译器
		verify.New()
		// 初始化异步任务
		asynq.New()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() {
		// 关闭数据库连接.
		dao.DBClose()
		// 输出协程数量.
		logger.Infof("[*]协程数量->%d\n", runtime.NumGoroutine())
		// 清除缓冲日志.
		logger.Sync()
	}()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// 获取 flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "dev", "config file (default is $HOME/.my_gin.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		panic("必须指定配制文件路径")
	}
	// 设置环境变量
	env.SetEnv(cfgFile)

	viper.SetConfigType("yml")
	// 读取 embed 方式 编译文件
	if err := viper.ReadConfig(config.DoConfig()); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Blue("no such config file")
		}
		panic(err)
	}
	viper.SetConfigName(env.Active().Value())
	// viper.AddConfigPath("/Users/xiaozhaofu/go/src/mygin/config/env")
	viper.AddConfigPath("/Users/xiaozhaofu/go/src/mygin/config/env")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("----config file changed:", e.String(), "---- config file name:", e.Name)
	})
}
