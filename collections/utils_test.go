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
