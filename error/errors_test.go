package error

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrNonNilDo(t *testing.T) {
	type args[T any] struct {
		one T
		err error
		f   func(e error)
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "success",
			args: args[int]{one: 1, err: nil, f: func(e error) {}},
			want: 1,
		},
		{
			name: "fail",
			args: args[int]{one: 1, err: errors.New("test"), f: func(e error) { panic(e) }},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("code not panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("expect error msg is test,but current is %v", r)
						return
					}
				}()
			}
			if got := ErrNonNilDo(tt.args.one, tt.args.err, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrNonNilDo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNonNilDo2(t *testing.T) {
	type args[T any, E any] struct {
		one T
		two E
		err error
		f   func(e error)
	}
	type testCase[T any, E any] struct {
		name  string
		args  args[T, E]
		want  T
		want1 E
	}
	tests := []testCase[string, int]{
		{
			name:  "success",
			args:  args[string, int]{"test", 1, nil, func(e error) {}},
			want:  "test",
			want1: 1,
		},
		{
			name:  "fail",
			args:  args[string, int]{"test", 1, errors.New("test"), func(e error) { panic(e) }},
			want:  "test",
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("code not panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("expect error msg is test,but current is %v", r)
						return
					}
				}()
			}
			got, got1 := ErrNonNilDo2(tt.args.one, tt.args.two, tt.args.err, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrNonNilDo2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ErrNonNilDo2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

}

func TestErrNonNilPanic(t *testing.T) {
	type args[T any] struct {
		one T
		err error
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "success",
			args: args[int]{one: 1, err: nil},
			want: 1,
		}, {
			name: "fail",
			args: args[int]{one: 1, err: errors.New("test")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("need panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("error msg must be test,but is %v", r)
						return
					}
				}()
			}
			if got := ErrNonNilPanic(tt.args.one, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrNonNilPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrNonNilPanic2(t *testing.T) {
	type args[T any, E any] struct {
		one T
		two E
		err error
	}
	type testCase[T any, E any] struct {
		name  string
		args  args[T, E]
		want  T
		want1 E
	}
	tests := []testCase[int, int]{
		{
			name:  "success",
			args:  args[int, int]{1, 2, nil},
			want:  1,
			want1: 2,
		},
		{
			name:  "fail",
			args:  args[int, int]{1, 2, errors.New("test")},
			want:  1,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("need panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("error msg must be test,but is %v", r)
						return
					}
				}()
			}
			got, got1 := ErrNonNilPanic2(tt.args.one, tt.args.two, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrNonNilPanic2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ErrNonNilPanic2() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIfErrNonNil(t *testing.T) {
	type args struct {
		err error
		f   func(e error)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				err: nil,
				f:   nil,
			},
		},
		{
			name: "fail",
			args: args{
				err: errors.New("test"),
				f: func(e error) {
					panic(e)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("need panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("error msg must be test,but is %v", r)
						return
					}
				}()
			}
			IfErrNonNil(tt.args.err, tt.args.f)
		})
	}
}

func TestIfErrNonNilPanic(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{err: nil},
		},
		{
			name: "fail",
			args: args{err: errors.New("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "fail" {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("need panic")
						return
					}
					if r != nil && r.(error).Error() != "test" {
						t.Errorf("error msg must be test,but is %v", r)
						return
					}
				}()
			}
			IfErrNonNilPanic(tt.args.err)
		})
	}
}
