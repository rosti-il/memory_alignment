package counter

import "sync"

type MutexCounter struct {
	value int32
	mutex sync.Mutex
}

func (c *MutexCounter) Inc(int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.value++
}

func (c *MutexCounter) Get() int32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.value
}
