package counter

type AlignedAtomicCounter struct {
	AtomicCounter
	_ [60]byte
}

type ShardedAlignedAtomicCounter struct {
	shards []AlignedAtomicCounter
}

func NewShardedAlignedAtomicCounter(shards int) ShardedAlignedAtomicCounter {
	return ShardedAlignedAtomicCounter{
		shards: make([]AlignedAtomicCounter, shards),
	}
}

func (c *ShardedAlignedAtomicCounter) Inc(idx int) {
	c.shards[idx].value.Add(1)
}

func (c *ShardedAlignedAtomicCounter) Get() int32 {
	var res int32

	for i := range c.shards {
		res = res + c.shards[i].value.Load()
	}

	return res
}
