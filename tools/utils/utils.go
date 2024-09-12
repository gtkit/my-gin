package utils

import (
	"context"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val any) bool {
	if val == nil {
		return true
	}
	return reflect.ValueOf(val).IsZero()
}

// IsNil The argument must be a chan, func, interface, map, pointer, or slice value;
// 判断一个对象是否真的是nil
func IsNil(val any) bool {
	if val == nil {
		return true
	}
	return reflect.ValueOf(val).IsNil()
}

// MillTime 获取当前时间戳(毫秒)
func MillTime() string {
	milltime := time.Now().UnixMilli()
	return strconv.FormatInt(milltime, 10)
}

// UnixTime 获取当前时间戳(秒)
func UnixTime() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

// Ternary 工具, 类似于 PHP 的三元运算符.
/**
func main() {
  fmt.Println(Ternary(true, 1, 2)) // 1
  fmt.Println(Ternary(false, 1, 2)) // 2
}
*/
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}

	return b
}

func Sleep(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// getAlloc 获取当前分配的内存
func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
