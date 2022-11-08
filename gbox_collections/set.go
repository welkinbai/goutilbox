package gbox_collections

type Set[E any] interface {
	Size() int
	IsEmpty() bool
	Contains(o E) bool
	Range(func(o E) bool)
	ToArray() []E
	Add(e E) bool
	Remove(o E) bool
	ContainsAll(c Set[E]) bool
	AddAll(c Set[E]) bool
	RetainAll(c Set[E]) bool
	RemoveAll(c Set[E]) bool
	Clear()
	Equals(c Set[E]) bool
}
