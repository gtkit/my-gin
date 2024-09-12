// @Author 2024/2/5 15:36:00
package test_test

import (
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	mrand2 "math/rand/v2"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	items := []int{2, 3, 4, 5}
	// 使用当前时间作为随机种子
	rng := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	// 生成一个介于0和切片长度之间的随机整数
	index := rng.Intn(len(items))
	t.Log("Random index: ", index)
	// 从切片中取出随机的元素
	item := items[index]
	t.Log("Random item: ", item)
}

func TestRandV2(t *testing.T) {
	items := []int{2, 3, 4, 5}
	l := len(items)

	i := mrand2.N(l)
	t.Log("max Int value: ", i)
	// item := items[i]
	// t.Log("Random item: ", item)
}

func TestCryptoRand(t *testing.T) {
	items := []int{2, 3, 4, 5}
	b := new(big.Int).SetInt64(int64(len(items)))
	t.Log("max Int value: ", b)
	i, err := rand.Int(rand.Reader, b)
	if err != nil {
		t.Error(err)
	}
	// 从切片中取出随机的元素
	index := i.Int64()
	t.Log("Random index: ", index)
	item := items[index]
	t.Log("Random item: ", item)
}

func TestCryptoRand2(t *testing.T) {
	// buf := make([]byte, 8)
	// _, err := rand.Read(buf)
	n, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("buf---%x", n)
}

type MySource struct {
}

func (m *MySource) Uint64() uint64 {
	return uint64(time.Now().UnixNano())
}
