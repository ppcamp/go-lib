package asyncheap

import (
	"container/heap"
	"sync"
)

// SafeHeap implement a concurrent safe heap.
type asyncHeap struct {
	locker sync.RWMutex
	h      heap.Interface
}

// Push x into heap.
func (h *asyncHeap) Push(x interface{}) {
	h.locker.Lock()
	defer h.locker.Unlock()
	heap.Push(h.h, x)
}

// Pop the minimum or maximum.
func (h *asyncHeap) Pop() interface{} {
	h.locker.Lock()
	defer h.locker.Unlock()
	return heap.Pop(h.h)
}

// Len get length.
func (h *asyncHeap) Len() int {
	h.locker.RLock()
	defer h.locker.RUnlock()
	return h.h.Len()
}

// New create new safe heap.
func New(impl heap.Interface) AsyncHeap {
	sh := &asyncHeap{h: impl}
	heap.Init(sh.h)
	return sh
}
