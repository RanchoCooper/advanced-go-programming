package ch01_basic

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
}

var (
	instance    *Singleton
	// by mutex
	initialized uint32
	mu          sync.Mutex
	// by once
	once sync.Once
)

func Instance() *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &Singleton{}
	}

	return instance
}

func InstanceByOnce() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
