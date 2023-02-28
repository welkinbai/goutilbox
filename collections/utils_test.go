package collections

import (
	"reflect"
	"testing"
)

func Test_findDupElement(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2 dup", args{[]int{1, 2, 3, 1}}, []int{1, 1}},
		{"3 dup", args{[]int{1, 2, 1, 3, 1}}, []int{1, 1, 1}},
		{"2 different dup", args{[]int{1, 2, 1, 3, 2}}, []int{1, 2, 1, 2}},
		{"0 dup", args{[]int{1, 2, 3, 4, 5}}, []int{}},
		{"all dup", args{[]int{1, 1, 1, 1, 1}}, []int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDupElement(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDupElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestOne struct {
	name string
	num  int
}

func Test_findDupElementInUnComparable(t *testing.T) {
	type args struct {
		slice              []TestOne
		compareElementFunc func(e TestOne) string
	}
	tests := []struct {
		name string
		args args
		want []TestOne
	}{
		{"2 dup", args{
			slice: []TestOne{{"just1", 1}, {"just2", 2}, {"just1", 3}},
			compareElementFunc: func(e TestOne) string {
				return e.name
			},
		}, []TestOne{{"just1", 1}, {"just1", 3}}},
		{"3 dup", args{
			slice: []TestOne{{"just1", 1}, {"just2", 2}, {"just1", 3}, {"just3", 4}, {"just1", 5}},
			compareElementFunc: func(e TestOne) string {
				return e.name
			},
		}, []TestOne{{"just1", 1}, {"just1", 3}, {"just1", 5}}},
		{"2 different dup", args{
			slice: []TestOne{{"just1", 1}, {"just2", 2}, {"just1", 3}, {"just3", 4}, {"just2", 5}},
			compareElementFunc: func(e TestOne) string {
				return e.name
			},
		}, []TestOne{{"just1", 1}, {"just2", 2}, {"just1", 3}, {"just2", 5}}},
		{"0 dup", args{
			slice: []TestOne{{"just1", 1}, {"just2", 2}, {"just3", 3}, {"just4", 4}, {"just5", 5}},
			compareElementFunc: func(e TestOne) string {
				return e.name
			},
		}, []TestOne{}},
		{"all dup", args{
			slice: []TestOne{{"just1", 1}, {"just1", 2}, {"just1", 3}, {"just1", 4}, {"just1", 5}},
			compareElementFunc: func(e TestOne) string {
				return e.name
			},
		}, []TestOne{{"just1", 1}, {"just1", 2}, {"just1", 3}, {"just1", 4}, {"just1", 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDupElementInUnComparable(tt.args.slice, tt.args.compareElementFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDupElementInUnComparable() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestObj struct {
	i   int
	j   int
	str string
}

func TestGroupListToMap(t *testing.T) {
	type args[T any, K comparable] struct {
		list    []T
		keyFunc func(t T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K][]T
	}
	tests := []testCase[TestObj, int]{
		{
			name: "success",
			args: args[TestObj, int]{list: []TestObj{{i: 1, j: 2}, {i: 1, j: 3}, {i: 2, j: 4}}, keyFunc: func(t TestObj) int { return t.i }},
			want: map[int][]TestObj{1: {{i: 1, j: 2}, {i: 1, j: 3}}, 2: {{i: 2, j: 4}}},
		},
		{
			name: "single_success",
			args: args[TestObj, int]{list: []TestObj{{i: 1, j: 2}}, keyFunc: func(t TestObj) int { return t.i }},
			want: map[int][]TestObj{1: {{i: 1, j: 2}}},
		},
		{
			name: "empty",
			args: args[TestObj, int]{list: []TestObj{}, keyFunc: func(t TestObj) int { return t.i }},
			want: map[int][]TestObj{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupListToMap(tt.args.list, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupListToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupListToMapString(t *testing.T) {
	type args[T any, K comparable] struct {
		list    []T
		keyFunc func(t T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K][]T
	}
	tests := []testCase[TestObj, string]{
		{
			name: "success",
			args: args[TestObj, string]{list: []TestObj{{i: 1, j: 2, str: "test1"}, {i: 1, j: 3, str: "test1"}, {i: 2, j: 4, str: "test2"}}, keyFunc: func(t TestObj) string { return t.str }},
			want: map[string][]TestObj{"test1": {{i: 1, j: 2, str: "test1"}, {i: 1, j: 3, str: "test1"}}, "test2": {{i: 2, j: 4, str: "test2"}}},
		},
		{
			name: "single_success",
			args: args[TestObj, string]{list: []TestObj{{i: 1, j: 2, str: "test1"}}, keyFunc: func(t TestObj) string { return t.str }},
			want: map[string][]TestObj{"test1": {{i: 1, j: 2, str: "test1"}}},
		},
		{
			name: "empty",
			args: args[TestObj, string]{list: []TestObj{}, keyFunc: func(t TestObj) string { return t.str }},
			want: map[string][]TestObj{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupListToMap(tt.args.list, tt.args.keyFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupListToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
