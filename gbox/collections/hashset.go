package collections

import "math"

type HashSet[E comparable] struct {
	m map[E]any
}

var dummyValue any = struct{}{}

func NewHashSet[E comparable]() *HashSet[E] {
	return &HashSet[E]{m: make(map[E]any, 16)}
}

func NewHashSetWith[E comparable](c HashSet[E]) *HashSet[E] {
	m := make(map[E]any, int(math.Max(float64(c.Size())/0.75+1.0, 16)))
	return &HashSet[E]{m: m}
}

func NewHashSetWithSize[E comparable](size int) *HashSet[E] {
	return &HashSet[E]{m: make(map[E]any, size)}
}

func (h HashSet[E]) Size() int {
	return len(h.m)
}

func (h HashSet[E]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h HashSet[E]) Contains(o E) bool {
	_, ok := h.m[o]
	return ok
}

func (h HashSet[E]) Range(f func(o E) bool) {
	for key := range h.m {
		if !f(key) {
			break
		}
	}
}

func (h HashSet[E]) ToArray() []E {
	result := make([]E, 0, len(h.m))
	for key := range h.m {
		result = append(result, key)
	}
	return result
}

func (h HashSet[E]) Add(e E) bool {
	_, ok := h.m[e]
	if ok {
		return false
	}
	h.m[e] = dummyValue
	return true
}

func (h HashSet[E]) Remove(o E) bool {
	_, ok := h.m[o]
	if !ok {
		return false
	}
	delete(h.m, o)
	return true
}

func (h HashSet[E]) ContainsAll(c Set[E]) bool {
	result := true
	c.Range(func(e E) bool {
		if !h.Contains(e) {
			result = false
			return false
		}
		return true
	})
	return result
}

func (h HashSet[E]) AddAll(c Set[E]) bool {
	modified := false
	c.Range(func(e E) bool {
		if h.Add(e) {
			modified = true
		}
		return true
	})
	return modified
}

func (h HashSet[E]) RetainAll(c Set[E]) bool {
	modified := false
	for key := range h.m {
		if !c.Contains(key) {
			delete(h.m, key)
			modified = true
		}
	}
	return modified
}

func (h HashSet[E]) RemoveAll(c Set[E]) bool {
	modified := false
	for key := range h.m {
		if c.Contains(key) {
			delete(h.m, key)
			modified = true
		}
	}
	return modified
}

func (h HashSet[E]) Clear() {
	for key := range h.m {
		delete(h.m, key)
	}
}

func (h HashSet[E]) Equals(c Set[E]) bool {
	if h.Size() != c.Size() {
		return false
	}
	return h.ContainsAll(c)
}
