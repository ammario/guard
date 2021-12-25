package guard

import "sync"

type Value[T any] struct {
	mu    sync.RWMutex
	inner T
}

// New creates a new type-guard.
func New[T any](v T) *Value[T] {
	return &Value[T]{
		inner: v,
	}
}

func (v *Value[T]) Write(f func(v T) T) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.inner = f(v.inner)
}

func (v *Value[T]) Read(f func(v T)) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	f(v.inner)
}
