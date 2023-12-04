package utils

import (
	"reflect"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	return reflect.ValueOf(val).IsZero()
}
