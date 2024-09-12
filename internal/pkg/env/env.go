package env

import (
	"strings"
)

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvPro  = "pro"
)

var (
	active Environment
	dev    Environment = &environment{value: EnvDev}
	test   Environment = &environment{value: EnvTest}
	pro    Environment = &environment{value: EnvPro}
)

var _ Environment = (*environment)(nil)

// Environment 环境配置.
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

// IsDev 是否是开发环境.
func (e *environment) IsDev() bool {
	return e.value == EnvDev
}

// IsTest 是否是测试环境.
func (e *environment) IsTest() bool {
	return e.value == EnvTest
}

// IsPro 是否是生产环境.
func (e *environment) IsPro() bool {
	return e.value == EnvPro
}

func (e *environment) t() {}

func SetEnv(env string) {
	switch strings.ToLower(strings.TrimSpace(env)) {
	case EnvDev:
		active = dev
	case EnvTest:
		active = test
	case EnvPro:
		active = pro
	default:
		active = dev
	}
}

// Active 当前配置的env.
func Active() Environment {
	return active
}

// IsDev 是否是开发环境.
func IsDev() bool {
	return active.IsDev()
}

// IsTest 是否是测试环境.
func IsTest() bool {
	return active.IsTest()
}

// IsPro 是否是生产环境.
func IsPro() bool {
	return active.IsPro()
}

func Env() string {
	return active.Value()
}

func Name() string {
	var envName string
	switch active.Value() {
	case EnvDev:
		envName = "⭐【开发环境】"
	case EnvTest:
		envName = "⭐【测试环境】"
	case EnvPro:
		envName = "🔴【生产环境】"
	default:
		envName = "【未知环境】"
	}
	return envName
}
