package jstypes

import (
	"reflect"
	"slices"
)

type SetForEachFunc[T any] func(element T, set *Set[T])

type Set[T any] struct {
	entries []T
}

func (set *Set[T]) Size() int {
	return len(set.entries)
}

func (set *Set[T]) Add(item T) *Set[T] {
	for _, i := range set.entries {
		if reflect.DeepEqual(item, i) {
			return set
		}
	}
	set.entries = append(set.entries, item)
	return set
}

func (set *Set[T]) Has(item T) bool {
	for _, i := range set.entries {
		if reflect.DeepEqual(item, i) {
			return true
		}
	}
	return false
}

func (set *Set[T]) Delete(item T) {
	var index *int
	for in, it := range set.entries {
		if reflect.DeepEqual(item, it) {
			index = &in
		}
	}
	if index == nil {
		return
	}
	set.entries = slices.Delete(set.entries, *index, *index+1)
}

func (set *Set[T]) Entries() []T {
	return set.entries
}

func (set *Set[T]) Clear() {
	set.entries = make([]T, 0)
}

func (set *Set[T]) ForEach(f SetForEachFunc[T]) {
	for _, v := range set.entries {
		f(v, set)
	}
}
