package cmd

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"ydsd_gin/config"
	"ydsd_gin/config/setup"
	"ydsd_gin/internal/dao"
	"ydsd_gin/internal/pkg/env"
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
		// 链接数据库
		dao.New()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	fmt.Println("----cmd init-----")
	cobra.OnInitialize(initConfig)

	// 获取 flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.ydsd_gin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("-------viper init with------ ", cfgFile)
	if cfgFile == "" {
		panic("必须指定配制文件路径")
	}
	// 设置环境变量
	env.SetEnv(cfgFile)
	fmt.Println("-------env value: ", env.Active().Value())

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
		fmt.Println("------redis addr", viper.GetString("redis.addr"))
		// 重新加载配制文件
	})

}
