package genetic

import (
	"reflect"
	"testing"
)

func Test_indexOf(t *testing.T) {
	type args[A comparable] struct {
		l []A
		a A
	}
	type testCase[A comparable] struct {
		name string
		args args[A]
		want int
	}
	tests := []testCase[int]{
		{"empty slice", args[int]{[]int{}, 42}, -1},
		{"singleton slice", args[int]{[]int{42}, 42}, 0},
		{"not in slice", args[int]{[]int{69}, 42}, -1},
		{"last in slice", args[int]{[]int{69, 420, 42}, 42}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOf(tt.args.l, tt.args.a); got != tt.want {
				t.Errorf("indexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	type args[A interface{}] struct {
		l []A
		i int
	}
	type testCase[A interface{}] struct {
		name     string
		args     args[A]
		wantS    []A
		wantElem A
		wantErr  bool
	}
	tests := []testCase[int]{
		{"empty slice", args[int]{[]int{}, 0}, []int{}, 0, true},
		{"singleton slice", args[int]{[]int{42}, 0}, []int{}, 42, false},
		{"index -1", args[int]{[]int{42}, -1}, []int{42}, 0, true},
		{"index 0", args[int]{[]int{42, 69}, 0}, []int{69}, 42, false},
		{"index len(l)-1", args[int]{[]int{42, 69}, 1}, []int{42}, 69, false},
		{"index len(l)", args[int]{[]int{42}, 1}, []int{42}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotElem, err := remove(tt.args.l, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("remove() gotS = %v, want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotElem, tt.wantElem) {
				t.Errorf("remove() gotElem = %v, want %v", gotElem, tt.wantElem)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args[A interface{}] struct {
		l []A
		a A
		i int
	}
	type testCase[A interface{}] struct {
		name      string
		args      args[A]
		wantSlice []A
		wantErr   bool
	}
	tests := []testCase[int]{
		{"index -1", args[int]{[]int{}, 42, -1}, []int{}, true},
		{"empty slice, index 0", args[int]{[]int{}, 42, 0}, []int{42}, false},
		{"index len(l)+1", args[int]{[]int{}, 42, 1}, []int{}, true},
		{"non-empty slice, index 0", args[int]{[]int{69}, 42, 0}, []int{42, 69}, false},
		{"non-empty slice, index len(l)", args[int]{[]int{69}, 42, 1}, []int{69, 42}, false},
		{"large slice, middle index", args[int]{[]int{42, 69, 420}, 5318008, 1}, []int{42, 5318008, 69, 420}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice, err := insert(tt.args.l, tt.args.a, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("insert() gotSlice = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
