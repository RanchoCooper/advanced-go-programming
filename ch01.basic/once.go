package ch01_basic

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	mu   sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
