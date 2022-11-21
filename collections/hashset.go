package collections

type hashSet[E comparable] struct {
	m map[E]any
}

var dummyValue any = struct{}{}

func NewHashSet[E comparable]() *hashSet[E] {
	return &hashSet[E]{m: make(map[E]any, 16)}
}

func NewHashSetWith[E comparable](c *hashSet[E]) *hashSet[E] {
	m := make(map[E]any, c.Size())
	newOne := hashSet[E]{m: m}
	newOne.AddAll(c)
	return &newOne
}

func NewHashSetWithSize[E comparable](size int) *hashSet[E] {
	return &hashSet[E]{m: make(map[E]any, size)}
}

func NewHashSetWithThese[E comparable](element ...E) *hashSet[E] {
	result := &hashSet[E]{m: make(map[E]any, len(element))}
	for _, ele := range element {
		result.Add(ele)
	}
	return result
}

func (h hashSet[E]) Size() int {
	return len(h.m)
}

func (h hashSet[E]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h hashSet[E]) Contains(o E) bool {
	_, ok := h.m[o]
	return ok
}

func (h hashSet[E]) Range(f func(o E) bool) {
	for key := range h.m {
		if !f(key) {
			break
		}
	}
}

func (h hashSet[E]) ToArray() []E {
	result := make([]E, 0, len(h.m))
	for key := range h.m {
		result = append(result, key)
	}
	return result
}

func (h hashSet[E]) Add(e E) bool {
	_, ok := h.m[e]
	if ok {
		return false
	}
	h.m[e] = dummyValue
	return true
}

func (h hashSet[E]) Remove(o E) bool {
	_, ok := h.m[o]
	if !ok {
		return false
	}
	delete(h.m, o)
	return true
}

func (h hashSet[E]) ContainsAll(c Set[E]) bool {
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

func (h hashSet[E]) AddAll(c Set[E]) bool {
	modified := false
	c.Range(func(e E) bool {
		if h.Add(e) {
			modified = true
		}
		return true
	})
	return modified
}

func (h hashSet[E]) RetainAll(c Set[E]) bool {
	modified := false
	for key := range h.m {
		if !c.Contains(key) {
			delete(h.m, key)
			modified = true
		}
	}
	return modified
}

func (h hashSet[E]) RemoveAll(c Set[E]) bool {
	modified := false
	for key := range h.m {
		if c.Contains(key) {
			delete(h.m, key)
			modified = true
		}
	}
	return modified
}

func (h hashSet[E]) Clear() {
	for key := range h.m {
		delete(h.m, key)
	}
}

func (h hashSet[E]) Equals(c Set[E]) bool {
	if h.Size() != c.Size() {
		return false
	}
	return h.ContainsAll(c)
}
