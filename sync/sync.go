package sync

import (
	"sync"

	"github.com/aszxqaz/automap"
)

type IMap[K comparable, V any] interface {
	Keys() []K
	Values() []V
	Len() int
	Get(k K) (V, bool)
	Set(k K, v V)
	Delete(k K) bool
	DeleteWhere(prd func(k K, v V) bool) bool
	Where(prd func(k K, v V) bool) (V, bool)
	Update(k K, fn func(k K, v V) V) bool
	UpdateWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool
	Reduce(init any, fn func(k K, v V, r any) any) any
	Transact(fn func(m automap.Map[K, V]))
}

type Map[K comparable, V any] struct {
	inner automap.Map[K, V]
	mu    sync.RWMutex
}

func (m *Map[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Keys()
}

func (m *Map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Values()
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Get(k)
}

func (m *Map[K, V]) Set(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.inner.Set(k, v)
}

func (m *Map[K, V]) Delete(k K) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inner.Delete(k)
}

func (m *Map[K, V]) DeleteWhere(prd func(k K, v V) bool) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inner.DeleteWhere(prd)
}

func (m *Map[K, V]) Where(prd func(k K, v V) bool) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Where(prd)
}

func (m *Map[K, V]) Update(k K, fn func(k K, v V) V) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inner.Update(k, fn)
}

func (m *Map[K, V]) UpdateWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inner.UpdateWhere(prd, fn)
}

func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Len()
}

func (m *Map[K, V]) Reduce(init any, fn func(k K, v V, r any) any) any {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.inner.Reduce(init, fn)
}

func (m *Map[K, V]) Transact(fn func(m automap.Map[K, V])) {
	m.mu.Lock()
	defer m.mu.Unlock()
	fn(m.inner)
}
