package collections

import (
	"reflect"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	type fields struct {
		m map[any]any
	}
	type args struct {
		e int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"addone", fields{}, args{e: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashSet[int]()
			if got := h.Add(tt.args.e); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestHashSet_AddAll(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		c Set
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.AddAll(tt.args.c); got != tt.want {
//				t.Errorf("AddAll() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_Clear(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			h.Clear()
//		})
//	}
//}
//
//func TestHashSet_Contains(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		o E
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.Contains(tt.args.o); got != tt.want {
//				t.Errorf("Contains() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_ContainsAll(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		c Set
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.ContainsAll(tt.args.c); got != tt.want {
//				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_Equals(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		c Set
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.Equals(tt.args.c); got != tt.want {
//				t.Errorf("Equals() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_IsEmpty(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.IsEmpty(); got != tt.want {
//				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_Range(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		f func(o E) bool
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			h.Range(tt.args.f)
//		})
//	}
//}
//
//func TestHashSet_Remove(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		o E
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.Remove(tt.args.o); got != tt.want {
//				t.Errorf("Remove() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_RemoveAll(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		c Set
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.RemoveAll(tt.args.c); got != tt.want {
//				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_RetainAll(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	type args struct {
//		c Set
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.RetainAll(tt.args.c); got != tt.want {
//				t.Errorf("RetainAll() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_Size(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.Size(); got != tt.want {
//				t.Errorf("Size() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestHashSet_ToArray(t *testing.T) {
//	type fields struct {
//		m map[any]any
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []E
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := HashSet{
//				m: tt.fields.m,
//			}
//			if got := h.ToArray(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ToArray() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
func TestNewHashSet(t *testing.T) {
	hashSet := NewHashSet[string]()
	tests := []struct {
		name string
		want HashSet[string]
	}{
		{"new", hashSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashSet; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestNewHashSetWith(t *testing.T) {
//	type args struct {
//		c HashSet
//	}
//	tests := []struct {
//		name string
//		args args
//		want HashSet
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewHashSetWith(tt.args.c); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewHashSetWith() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewHashSetWithSize(t *testing.T) {
//	type args struct {
//		size int
//	}
//	tests := []struct {
//		name string
//		args args
//		want HashSet
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewHashSetWithSize(tt.args.size); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewHashSetWithSize() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
