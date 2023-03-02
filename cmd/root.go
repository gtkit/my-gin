package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.superjq.com/go-tools/verify"

	"ydsd_gin/config"
	"ydsd_gin/config/setup"
	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/pkg/env"
	jwt "ydsd_gin/tools/jwtauth"
	"ydsd_gin/tools/logger"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 配置初始化
		setup.Initialize()
		// 初始化 logger
		logger.Init()
		// 初始化 jwt
		jwt.InitJwt()
		// 链接数据库, redis
		dao.New()
		// 初始化验证翻译器
		verify.New()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// 获取 flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "dev", "config file (default is $HOME/.ydsd_gin.yaml)")
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
		panic(err)
	}
	viper.SetConfigName(cfgFile)
	viper.AddConfigPath("./config/env")

	// 读取配制文件
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	// }

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Op, e.Name)
		// 重新加载配制文件
		setup.Initialize()
	})

}
