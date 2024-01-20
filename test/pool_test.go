package test_test

import (
	"fmt"
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

// 没有使用Sync.Pool的
func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			// p.Age = 23
		}
	}
}

// 带有Sync.Pool的对象
func BenchmarkWithPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p, ok := personPool.Get().(*Person)
			if !ok {
				fmt.Println("get error")
			}
			p.Age = 23
			fmt.Println(p)
			personPool.Put(p)
		}
	}
}
