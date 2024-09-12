// @Author 2024/2/5 15:41:00
package utils

import (
	"math/rand"
	"time"
)

type RandomTyper interface {
	int | int8 | int16 | int32 | int64 | string
}

// RandomChoice 切片中随机取一个元素
func RandomChoice[T RandomTyper](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}
	// 使用当前时间作为随机种子
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个介于0和切片长度之间的随机整数
	index := rng.Intn(len(slice))
	// 从切片中取出随机的元素
	return slice[index]
}
