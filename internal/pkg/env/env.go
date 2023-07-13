package env

import (
	"fmt"
	"strings"
)

var (
	active Environment
	dev    Environment = &environment{value: "dev"}
	test   Environment = &environment{value: "test"}
	pro    Environment = &environment{value: "pro"}
)

var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsPro() bool
	t()
}

type environment struct {
	value string
}

func (e *environment) Value() string {
	return e.value
}

// IsDev 是否是开发环境
func (e *environment) IsDev() bool {
	return e.value == "dev"
}

// IsTest 是否是测试环境
func (e *environment) IsTest() bool {
	return e.value == "test"
}

// IsPro 是否是生产环境
func (e *environment) IsPro() bool {
	return e.value == "pro"
}

func (e *environment) t() {}

func SetEnv(env string) {

	switch strings.ToLower(strings.TrimSpace(env)) {
	case "dev":
		active = dev
	case "test":
		active = test
	case "pro":
		active = pro
	default:
		active = dev
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default 'dev' will be used.")
	}
}

// Active 当前配置的env
func Active() Environment {
	return active
}

// IsDev 是否是开发环境
func IsDev() bool {
	return active.IsDev()
}

// IsTest 是否是测试环境
func IsTest() bool {
	return active.IsTest()
}

// IsPro 是否是生产环境
func IsPro() bool {
	return active.IsPro()
}
