package automap

type Map[K comparable, V any] struct {
	inner map[K]V
}

func (m *Map[K, V]) Keys() []K {
	m.maybeInit()
	keys := make([]K, 0, len(m.inner))
	for k := range m.inner {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	m.maybeInit()
	values := make([]V, 0, len(m.inner))
	for _, v := range m.inner {
		values = append(values, v)
	}
	return values
}

func (m *Map[K, V]) Len() int {
	m.maybeInit()
	return len(m.inner)
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

func (m *Map[K, V]) DeleteWhere(prd func(k K, v V) bool) bool {
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

func (m *Map[K, V]) FirstWhere(prd func(k K, v V) bool) (V, bool) {
	m.maybeInit()
	for k, v := range m.inner {
		if prd(k, v) {
			return v, true
		}
	}
	var v V
	return v, false
}

func (m *Map[K, V]) ValuesWhere(prd func(k K, v V) bool) []V {
	m.maybeInit()
	vs := []V{}
	for k, v := range m.inner {
		if prd(k, v) {
			vs = append(vs, v)
		}
	}
	return vs
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
