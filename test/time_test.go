package test_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub
	// 将t0加d获取t1就是使用Add方法
	k := time.Now()
	// 一天之前
	d, _ := time.ParseDuration("-24h")
	fmt.Println(k.Add(d))
	// 一周之前
	fmt.Println(k.Add(d * 7))
	// 一月之前
	fmt.Println(k.Add(d * 30))
}

func TestMake(t *testing.T) {
	a := make([]int, 3, 10) // 切片长度为 1，预留空间长度为 10
	a = append(a, 1)
	fmt.Printf("%v--%T ---%d\n", a, a, len(a))
}
