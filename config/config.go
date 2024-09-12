// @Author xiaozhaofu 2023/2/19 18:01:00
package config

import (
	"bytes"
	_ "embed"
	"io"

	"github.com/json-iterator/go/extra"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"ydsd_gin/internal/pkg/env"
	"ydsd_gin/tools/utils"
)

var (
	//go:embed env/dev.yml
	devConfigs []byte

	//go:embed env/test.yml
	testConfigs []byte

	//go:embed env/pro.yml
	proConfigs []byte
)

func init() {
	ConfigFuncs = make(map[string]ConfigFunc)
	// json 兼容 PHP 数字字符串格式
	extra.RegisterFuzzyDecoders()
	// fmt.Println("----set json extra.RegisterFuzzyDecoders ------")
}

func DoConfig() (r io.Reader) {
	switch env.Active().Value() {
	case "dev":
		r = bytes.NewReader(devConfigs)
	case "test":
		r = bytes.NewReader(testConfigs)
	case "pro":
		r = bytes.NewReader(proConfigs)
	default:
		r = bytes.NewReader(devConfigs)
	}

	return
}

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]any

// ConfigFuncs 先加载到此数组，loadConfig 在动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func LoadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...any) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...any) any {
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || utils.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...any) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...any) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...any) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}

func GetStringMapStringSlice(path string) map[string][]string {
	return viper.GetStringMapStringSlice(path)
}
func GetStringSlice(path string) []string {
	return viper.GetStringSlice(path)
}

func GetIntSlice(path string) []int {
	return viper.GetIntSlice(path)
}
