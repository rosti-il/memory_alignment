package counter

import (
	"runtime"
	"sync"
	"testing"

	c "example.com/demo/internal/counter"
)

func BenchmarkRWLocked(b *testing.B) {
	var counter c.RWMutexCounter
	work(b, &counter)
}

func BenchmarkLocked(b *testing.B) {
	var counter c.MutexCounter
	work(b, &counter)
}

func BenchmarkAtomic(b *testing.B) {
	var counter c.AtomicCounter
	work(b, &counter)
}

func BenchmarkShardedAtomic(b *testing.B) {
	counter := c.NewShardedAtomicCounter(runtime.NumCPU())
	work(b, &counter)
}

func BenchmarkShardedAlignedAtomic(b *testing.B) {
	counter := c.NewShardedAlignedAtomicCounter(runtime.NumCPU())
	work(b, &counter)
}

func work(b *testing.B, counter Counter) {
	var wg sync.WaitGroup

	wg.Add(runtime.NumCPU())
	for cpuid := range runtime.NumCPU() {
		go func(id int) {
			defer wg.Done()
			for range b.N {
				for range 10000 {
					counter.Inc(id)
				}
				counter.Get()
			}
		}(cpuid)
	}

	wg.Wait()
}
