package jstypes

import (
	"fmt"
	"math"
	"reflect"
	"slices"
)

type ArrayFilterFunc[T any] func(element T, index int, array *Array[T]) bool
type ArrayForEachFunc[T any] func(element T, index int, array *Array[T])
type ArrayMapFunc[T any] func(element T, index int, array *Array[T]) T

type Array[T any] []T

func (a *Array[T]) Push(items ...T) {
	*a = append(*a, items...)
}

func (a *Array[T]) At(index int) *T {
	if index < 0 {
		b := make(Array[T], len(*a))
		copy(b, *a)
		b.Reverse()
		index = int(math.Abs(float64(index))) - 1
		if len(b) <= index {
			return nil
		}
		return &b[index]
	} else {
		if len(*a) <= index {
			return nil
		}
		return &(*a)[index]
	}
}

func (a *Array[T]) Concat(b Array[T]) Array[T] {
	n := make(Array[T], len(*a))
	copy(n, *a)
	for _, i := range b {
		n.Push(i)
	}
	return n
}

func (a *Array[T]) Filter(f ArrayFilterFunc[T]) Array[T] {
	n := make(Array[T], 0)
	for in, it := range *a {
		if f(it, in, a) {
			n.Push(it)
		}
	}
	return n
}

func (a *Array[T]) Every(f ArrayFilterFunc[T]) bool {
	for in, it := range *a {
		if !f(it, in, a) {
			return false
		}
	}
	return true
}

func (a *Array[T]) Some(f ArrayFilterFunc[T]) bool {
	for in, it := range *a {
		if f(it, in, a) {
			return true
		}
	}
	return false
}

func (a *Array[T]) Find(f ArrayFilterFunc[T]) *T {
	for in, it := range *a {
		if f(it, in, a) {
			return &it
		}
	}
	return nil
}

func (a *Array[T]) FindLast(f ArrayFilterFunc[T]) *T {
	b := make(Array[T], len(*a))
	copy(b, *a)
	b.Reverse()
	for in, it := range b {
		if f(it, in, a) {
			return &it
		}
	}
	return nil
}

func (a *Array[T]) Length() int {
	return len(*a)
}

func (a *Array[T]) ForEach(f ArrayForEachFunc[T]) {
	for in, it := range *a {
		f(it, in, a)
	}
}

func (a *Array[T]) IndexOf(element T) int {
	for i, e := range *a {
		if reflect.DeepEqual(e, element) {
			return i
		}
	}
	return -1
}

func (a *Array[T]) Map(f ArrayMapFunc[T]) Array[T] {
	n := make(Array[T], 0)
	for in, it := range *a {
		n.Push(f(it, in, a))
	}
	return n
}

func (a *Array[T]) Reverse() {
	slices.Reverse(*a)
}

func (a *Array[T]) Slice(start int, e ...int) Array[T] {
	end := len(*a)
	if len(e) == 1 {
		end = e[0]
	}
	return (*a)[start:end]
}

func (a *Array[T]) Pop() *T {
	e := (*a)[a.Length()-1]
	*a = slices.Delete(*a, a.Length()-1, a.Length())
	return &e
}

func (a *Array[T]) Shift() *T {
	if len(*a) == 0 {
		return nil
	}
	e := (*a)[0]
	*a = slices.Delete(*a, 0, 1)
	return &e
}

func (a *Array[T]) Join(joiner string) (str string) {
	for in, it := range *a {
		str += fmt.Sprint(it)
		if in < len(*a)-1 {
			str += joiner
		}
	}
	return
}

func (a *Array[T]) ToString() string {
	return a.Join(",")
}

func (a *Array[T]) Unshift(items ...T) {
	b := make(Array[T], len(*a))
	copy(b, *a)
	b.Reverse()
	slices.Reverse(items)
	b = append(b, items...)
	b.Reverse()
	*a = b
}

func (a *Array[T]) Includes(element T) bool {
	for _, e := range *a {
		if reflect.DeepEqual(element, e) {
			return true
		}
	}
	return false
}

func (a *Array[T]) With(index int, element T) Array[T] {
	b := make(Array[T], len(*a))
	copy(b, *a)
	if b.Length() <= index {
		a.Push(element)
	}
	b[index] = element
	return b
}

func (a *Array[T]) ToSlice() []T {
	return []T(*a)
}
