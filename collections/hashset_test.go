package collections

import (
	"reflect"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	type fields struct {
		m map[int]any
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
		{"addOne", fields{m: make(map[int]any, 16)}, args{e: 3}, true},
		{"addAlreadyExist", fields{m: map[int]any{3: struct{}{}}}, args{e: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.Add(tt.args.e); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_AddAll(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		c Set[int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "addAll",
			fields: fields{m: make(map[int]any, 16)},
			args:   args{c: hashSet[int]{m: map[int]any{1: struct{}{}, 2: struct{}{}}}},
			want:   true,
		},
		{
			name:   "addAllExist",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:   args{c: hashSet[int]{m: map[int]any{1: struct{}{}, 2: struct{}{}}}},
			want:   false,
		},
		{
			name:   "addAllEmpty",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:   args{c: hashSet[int]{m: map[int]any{}}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.AddAll(tt.args.c); got != tt.want {
				t.Errorf("AddAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Clear(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "clear",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			h.Clear()
			if len(h.m) != 0 {
				t.Errorf("after Clear(),map size not be 0")
			}
		})
	}
}

func TestHashSet_Contains(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		o int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "containsYes",
			fields: fields{m: map[int]any{1: struct{}{}}},
			args:   args{o: 1},
			want:   true,
		},
		{
			name:   "containsNo",
			fields: fields{m: map[int]any{1: struct{}{}}},
			args:   args{o: 3},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.Contains(tt.args.o); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_ContainsAll(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		c Set[int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "containsAllYes",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:   args{c: NewHashSetWithThese(1, 2)},
			want:   true,
		},
		{
			name:   "containsAllNo",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:   args{c: NewHashSetWithThese(1, 2, 3)},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.ContainsAll(tt.args.c); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Equals(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		c Set[int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "equalsYes",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			args:   args{c: hashSet[int]{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}}},
			want:   true,
		},
		{
			name:   "equalsNo",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			args:   args{c: hashSet[int]{m: map[int]any{1: struct{}{}, 2: struct{}{}}}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.Equals(tt.args.c); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "isEmpty_Yes",
			fields: fields{m: map[int]any{}},
			want:   true,
		},
		{
			name:   "isEmpty_No",
			fields: fields{m: map[int]any{1: struct{}{}}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Range(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		f func(o int) bool
	}
	var collector []int
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "rangeAll",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			args: args{f: func(o int) bool {
				collector = append(collector, o)
				return true
			}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			h.Range(tt.args.f)
			if !reflect.DeepEqual(collector, tt.want) {
				t.Errorf("Range = %v, want %v", collector, tt.want)
			}
		})
	}
}

func TestHashSet_Remove(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		o int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       bool
		afterCheck map[int]any
	}{
		{
			name:       "remove_true",
			fields:     fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:       args{1},
			want:       true,
			afterCheck: map[int]any{2: struct{}{}},
		},
		{
			name:       "remove_false",
			fields:     fields{m: map[int]any{1: struct{}{}, 2: struct{}{}}},
			args:       args{3},
			want:       false,
			afterCheck: map[int]any{1: struct{}{}, 2: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.Remove(tt.args.o); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(h.m, tt.afterCheck) {
				t.Errorf("after Remove() map:%v, want %v", h.m, tt.afterCheck)
			}
		})
	}
}

func TestHashSet_RemoveAll(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		c Set[int]
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       bool
		afterCheck map[int]any
	}{
		{
			name: "removeAll_true",
			fields: fields{
				m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			},
			args:       args{c: NewHashSetWithThese(1, 2)},
			want:       true,
			afterCheck: map[int]any{3: struct{}{}},
		},
		{
			name: "removeAll_false",
			fields: fields{
				m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			},
			args:       args{c: NewHashSetWithThese(4, 5)},
			want:       false,
			afterCheck: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.RemoveAll(tt.args.c); got != tt.want {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(h.m, tt.afterCheck) {
				t.Errorf("after RemoveAll() map:%v, want %v", h.m, tt.afterCheck)
			}
		})
	}
}

func TestHashSet_RetainAll(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	type args struct {
		c Set[int]
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       bool
		afterCheck map[int]any
	}{
		{
			name:       "retainAll_true",
			fields:     fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			args:       args{c: NewHashSetWithThese(2, 3)},
			want:       true,
			afterCheck: map[int]any{2: struct{}{}, 3: struct{}{}},
		},
		{
			name:       "retainAll_false",
			fields:     fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			args:       args{c: NewHashSetWithThese(1, 2, 3)},
			want:       false,
			afterCheck: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.RetainAll(tt.args.c); got != tt.want {
				t.Errorf("RetainAll() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(h.m, tt.afterCheck) {
				t.Errorf("after RetainAll() map:%v, want %v", h.m, tt.afterCheck)
			}
		})
	}
}

func TestHashSet_Size(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "size",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			want:   3,
		},
		{
			name:   "emptySize",
			fields: fields{m: map[int]any{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_ToArray(t *testing.T) {
	type fields struct {
		m map[int]any
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "toArray",
			fields: fields{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
			want:   []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashSet[int]{
				m: tt.fields.m,
			}
			got := h.ToArray()
			for _, w := range tt.want {
				exist := false
				for _, g := range got {
					if g == w {
						exist = true
					}
				}
				if !exist {
					t.Errorf("ToArray() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestNewHashSet(t *testing.T) {
	tests := []struct {
		name string
		want *hashSet[string]
	}{
		{"new", &hashSet[string]{
			m: make(map[string]any, 16),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashSet[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashSetWith(t *testing.T) {
	type args struct {
		c *hashSet[string]
	}
	tests := []struct {
		name string
		args args
		want *hashSet[string]
	}{
		{
			name: "new",
			args: args{c: &hashSet[string]{m: map[string]any{
				"test1": 1,
			}}},
			want: &hashSet[string]{m: map[string]any{"test1": struct{}{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashSetWith(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashSetWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashSetWithSize(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *hashSet[string]
	}{
		{
			name: "new",
			args: args{size: 100},
			want: &hashSet[string]{m: make(map[string]any, 100)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashSetWithSize[string](tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashSetWithSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHashSetWithThese(t *testing.T) {
	type args struct {
		element []int
	}
	tests := []struct {
		name string
		args args
		want *hashSet[int]
	}{
		{
			name: "new",
			args: args{[]int{1, 2, 3}},
			want: &hashSet[int]{m: map[int]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashSetWithThese(tt.args.element...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashSetWithThese() = %v, want %v", got, tt.want)
			}
		})
	}
}
