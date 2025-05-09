package counter

import "sync/atomic"

type AtomicCounter struct {
	value atomic.Int32
}

func (c *AtomicCounter) Inc(int) {
	c.value.Add(1)
}

func (c *AtomicCounter) Get() int32 {
	return c.value.Load()
}
