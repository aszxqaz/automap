package automap

type IMap[K comparable, V any] interface {
	Get(k K) (V, bool)
	Set(k K, v V)
	Delete(k K) bool
	DeleteWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool
	Where(prd func(k K, v V) bool) (V, bool)
	Update(k K, fn func(k K, v V) V) bool
	UpdateWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool
	Len() int
	Reduce(init any, fn func(k K, v V, r any) any) any
}

type Map[K comparable, V any] struct {
	inner map[K]V
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	m.maybeInit()
	v, ok := m.inner[k]
	return v, ok
}

func (m *Map[K, V]) Set(k K, v V) {
	m.maybeInit()
	m.inner[k] = v
}

func (m *Map[K, V]) Delete(k K) bool {
	m.maybeInit()
	_, ok := m.inner[k]
	if ok {
		delete(m.inner, k)
		return true
	}
	return false
}

func (m *Map[K, V]) DeleteWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool {
	m.maybeInit()
	ok := false
	for k, v := range m.inner {
		if prd(k, v) {
			delete(m.inner, k)
			ok = true
		}
	}
	return ok
}

func (m *Map[K, V]) Where(prd func(k K, v V) bool) (V, bool) {
	m.maybeInit()
	for k, v := range m.inner {
		if prd(k, v) {
			return v, true
		}
	}
	var v V
	return v, false
}

func (m *Map[K, V]) Update(k K, fn func(k K, v V) V) bool {
	m.maybeInit()
	v, ok := m.inner[k]
	if ok {
		m.inner[k] = fn(k, v)
		return true
	}
	return false
}

func (m *Map[K, V]) UpdateWhere(prd func(k K, v V) bool, fn func(k K, v V) V) bool {
	m.maybeInit()
	ok := false
	for ik, iv := range m.inner {
		if prd(ik, iv) {
			m.inner[ik] = fn(ik, iv)
			ok = true
		}
	}
	return ok
}

func (m *Map[K, V]) Len() int {
	m.maybeInit()
	return len(m.inner)
}

func (m *Map[K, V]) Reduce(init any, fn func(k K, v V, r any) any) any {
	m.maybeInit()
	res := init
	for ik, iv := range m.inner {
		res = fn(ik, iv, res)
	}
	return res
}

func (m *Map[K, V]) maybeInit() {
	if m.inner == nil {
		m.inner = make(map[K]V)
	}
}
