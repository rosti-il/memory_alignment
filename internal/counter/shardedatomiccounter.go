package counter

type ShardedAtomicCounter struct {
	shards []AtomicCounter
}

func NewShardedAtomicCounter(shards int) ShardedAtomicCounter {
	return ShardedAtomicCounter{
		shards: make([]AtomicCounter, shards),
	}
}

func (c *ShardedAtomicCounter) Inc(idx int) {
	c.shards[idx].value.Add(1)
}

func (c *ShardedAtomicCounter) Get() int32 {
	var res int32

	for i := range c.shards {
		res = res + c.shards[i].value.Load()
	}

	return res
}
