package tasks

import (
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	counter int64
}

type MutexCounter struct {
	mu    sync.Mutex
	value int
}

func (m *MutexCounter) Inc(delta int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value += delta
}

func (m *MutexCounter) Get() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.value
}

func (a *AtomicCounter) Inc(delta int64) {
	atomic.AndInt64(&a.counter, delta)
}

func (a *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&a.counter)
}

func (a *AtomicCounter) AtomicRepeat(n int) {
	for i := 0; i < n; i++ {
		atomic.AddInt64(&a.counter, 1)
	}
	wg.Done()
}

func (m *MutexCounter) MutexRepeat(n int) {
	for i := 0; i < n; i++ {
		m.Inc(1)
	}
	wg.Done()
}

var wg sync.WaitGroup

func RunAtomicCounter(incCount, goroutineCount int) int64 {
	var atomicCounter AtomicCounter
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go atomicCounter.AtomicRepeat(incCount)
	}
	wg.Wait()
	return atomicCounter.Get()
}

func RunMutexCounter(incCount, goroutineCount int) int {
	var mutexCounter MutexCounter
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go mutexCounter.MutexRepeat(incCount)
	}
	wg.Wait()
	return mutexCounter.Get()
}
