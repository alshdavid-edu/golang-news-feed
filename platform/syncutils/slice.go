package syncutils

import "sync"

type Slice[T any] struct {
	mutex sync.RWMutex
	items []T
}

type SliceItem[T any] struct {
	Index int
	Value T
}

func NewConcurrentSlice[T any]() *Slice[T] {
	cs := &Slice[T]{
		items: []T{},
	}

	return cs
}

func (cs *Slice[T]) Append(item ...T) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.items = append(cs.items, item...)
}

func (cs *Slice[T]) Items() []T {
	var items []T
	return append(items, cs.items...)
}

func (cs *Slice[T]) Len() int {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()

	return len(cs.items)
}
