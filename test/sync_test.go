package test

//
import (
	"sync"
	"testing"
)

func Benchmark_3(b *testing.B) {

	// 对比，分配一个大堆时，采用pool和不采用pool的性能对比
	// 协程数：100，每次需求对象大小1m
	const routineCount = 10
	const size = 1 << 20

	b.Run("no-pool", func(b *testing.B) {

		wg := sync.WaitGroup{}
		for i := 0; i < routineCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N; j++ {
					// 申请内存
					nopool := make([]byte, size)
					nopool[0] = 1
				}
			}()
		}
		wg.Wait()
	})

	b.Run("pool", func(b *testing.B) {
		wg := sync.WaitGroup{}
		pool := sync.Pool{
			New: func() interface{} {
				p := make([]byte, size)
				return p
			},
		}
		for i := 0; i < routineCount; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N; j++ {
					// 申请内存
					bpool := pool.Get().([]byte)
					bpool[0] = 1
					pool.Put(bpool)
				}
			}()
		}
		wg.Wait()
	})
}
