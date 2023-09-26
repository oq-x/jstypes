package jstypes

type MapForEachFunc[K comparable, V any] func(value V, key K, m *Map[K, V])

type Map[K comparable, V any] struct {
	m map[K]V
}

func MapFrom[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V]{m: m}
}

func (m *Map[K, V]) Get(key K) *V {
	if m.m == nil {
		m.m = make(map[K]V)
	}
	if i, ok := m.m[key]; !ok {
		return nil
	} else {
		return &i
	}
}

func (m *Map[K, V]) Set(key K, value V) {
	if m.m == nil {
		m.m = make(map[K]V)
	}
	m.m[key] = value
}

func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *Map[K, V]) ForEach(f MapForEachFunc[K, V]) {
	for k, v := range m.m {
		f(v, k, m)
	}
}

func (m *Map[K, V]) Delete(key K) {
	delete(m.m, key)
}

func (m *Map[K, V]) Clear() {
	m.m = make(map[K]V)
}

func (m *Map[K, V]) Keys() (arr Array[K]) {
	for k := range m.m {
		arr.Push(k)
	}
	return
}

func (m *Map[K, V]) Values() (arr Array[V]) {
	for _, v := range m.m {
		arr.Push(v)
	}
	return
}

func (m *Map[K, V]) Entries() (arr Array[Array[interface{}]]) {
	for k, v := range m.m {
		arr.Push(Array[interface{}]{
			k, v,
		})
	}
	return
}

func (m *Map[K, V]) ToMap() map[K]V {
	return m.m
}
