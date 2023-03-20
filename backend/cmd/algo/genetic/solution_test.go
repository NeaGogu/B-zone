package genetic

import (
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
	"math/rand"
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

func Test_dist(t *testing.T) {
	type args struct {
		p0 Pos
		p1 Pos
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"same", args{Pos{0, 0, 0}, Pos{0, 0, 0}}, 0},
		{"y+", args{Pos{0, 0, 0}, Pos{1, 0, 0}}, 1},
		{"y-", args{Pos{0, 0, 0}, Pos{-1, 0, 0}}, 1},
		{"x+", args{Pos{0, 0, 0}, Pos{0, 1, 0}}, 1},
		{"x-", args{Pos{0, 0, 0}, Pos{0, -1, 0}}, 1},
		{"diagonal45deg", args{Pos{0, 0, 0}, Pos{1, 1, 0}}, math.Sqrt2},
		{"diagonal_345", args{Pos{0, 0, 0}, Pos{3, 4, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dist(tt.args.p0, tt.args.p1); got != tt.want {
				t.Errorf("dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greedyRoute(t *testing.T) {
	type args struct {
		depot  Pos
		points []Pos
	}
	tests := []struct {
		name string
		args args
		want Route
	}{
		{"no points", args{Pos{0, 0, 0}, []Pos{}},
			Route{Pos{0, 0, 0}, []Pos{}}},
		{"1 points", args{Pos{0, 0, 0}, []Pos{{1, 1, 0}}},
			Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"2 equal points", args{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {1, 1, 0}}},
			Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {1, 1, 0}}}},
		{"2 different points ordered", args{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
			Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}},
		{"2 different points unordered", args{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}},
			Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greedyRoute(tt.args.depot, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("greedyRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoute_applySwap(t *testing.T) {
	tests := []struct {
		name    string
		route   Route
		i       int
		j       int
		wantErr bool
		want    Route
	}{
		{"empty route", Route{Pos{0, 0, 0}, []Pos{}}, 0, 0,
			true, Route{Pos{0, 0, 0}, []Pos{}}},
		{"singleton route", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 0,
			false, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"index to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative index", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"swap 2", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}, 0, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.applySwap(tt.i, tt.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("applySwap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.route, tt.want) {
				t.Errorf("applySwap() route = %v, want %v", tt.route, tt.want)
			}
		})
	}
}

func TestRoute_apply2Opt(t *testing.T) {
	tests := []struct {
		name    string
		route   Route
		i       int
		j       int
		wantErr bool
		want    Route
	}{
		{"empty route", Route{Pos{0, 0, 0}, []Pos{}}, 0, 0,
			true, Route{Pos{0, 0, 0}, []Pos{}}},
		{"singleton route", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 0,
			false, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"index to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative index", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"2opt 2", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}, 0, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}}},
		{"2opt many", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}, {4, 4, 0}, {5, 5, 0}}}, 1, 3,
			false, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {4, 4, 0}, {3, 3, 0}, {2, 2, 0}, {5, 5, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.apply2Opt(tt.i, tt.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("apply2Opt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.route, tt.want) {
				t.Errorf("apply2Opt() route = %v, want %v", tt.route, tt.want)
			}
		})
	}
}

func TestRoute_applyGreedy(t *testing.T) {
	tests := []struct {
		name    string
		route   Route
		i       int
		wantErr bool
		want    Route
	}{
		{"empty route", Route{Pos{0, 0, 0}, []Pos{}}, 0,
			true, Route{Pos{0, 0, 0}, []Pos{}}},
		{"singleton route", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0,
			false, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"optimal route", Route{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
		{"non-optimal route", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {0, 1, 0}, {1, 0, 0}}}, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
		{"index to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative index", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.applyGreedy(tt.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyGreedy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.route, tt.want) {
				t.Errorf("applyGreedy() route = %v, want %v", tt.route, tt.want)
			}
		})
	}
}

func TestSolution_applySwap(t *testing.T) {
	type args struct {
		r0 int
		i  int
		r1 int
		j  int
	}
	tests := []struct {
		name    string
		sol     Solution
		args    args
		wantErr bool
		want    Solution
	}{
		{
			name: "r0<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{-1, 0, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "i<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, -1, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r1<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, -1, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "j<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 0, -1},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r0>=len(sol.routes)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{1, 0, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "i>=len(sol.routes[r0].activities)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 1, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r1>=len(sol.routes)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 1, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "j>=len(sol.routes[r1].activities)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 0, 1},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "swap different route",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 1, 1, 0},
			wantErr: false,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {3, 3, 0}}},
					{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {4, 4, 0}}},
				},
				Cost: 0},
		},
		{
			name: "swap same route",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 1, 0, 0},
			wantErr: false,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sol.applySwap(tt.args.r0, tt.args.i, tt.args.r1, tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("applySwap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.sol, tt.want) {
				t.Errorf("applySwap() sol = %v, want %v", tt.sol, tt.want)
			}
		})
	}
}

func TestSolution_applyMigrate(t *testing.T) {
	type args struct {
		r0 int
		i  int
		r1 int
		j  int
	}
	tests := []struct {
		name    string
		sol     Solution
		args    args
		wantErr bool
		want    Solution
	}{
		{
			name: "r0<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{-1, 0, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "i<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, -1, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r1<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, -1, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "j<0",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 0, -1},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r0>=len(sol.routes)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{1, 0, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "i>=len(sol.routes[r0].activities)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 1, 0, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "r1>=len(sol.routes)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 1, 0},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "j>len(sol.routes[r1].activities)",
			sol: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
			args:    args{0, 0, 0, 2},
			wantErr: true,
			want: Solution{
				Routes: []Route{{
					Depot:      Pos{0, 0, 0},
					Activities: []Pos{{1, 1, 0}},
				}},
				Cost: 0},
		},
		{
			name: "same route",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 0, 0, 1},
			wantErr: true,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
		},
		{
			name: "migrate front",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 0, 1, 0},
			wantErr: false,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
		},
		{
			name: "migrate middle",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 0, 1, 1},
			wantErr: false,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {1, 1, 0}, {4, 4, 0}}},
				},
				Cost: 0},
		},
		{
			name: "migrate end",
			sol: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}}},
				},
				Cost: 0},
			args:    args{0, 0, 1, 2},
			wantErr: false,
			want: Solution{
				Routes: []Route{
					{Pos{0, 0, 0}, []Pos{{2, 2, 0}}},
					{Pos{0, 0, 0}, []Pos{{3, 3, 0}, {4, 4, 0}, {1, 1, 0}}},
				},
				Cost: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sol.applyMigrate(tt.args.r0, tt.args.i, tt.args.r1, tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("applyMigrate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.sol, tt.want) {
				t.Errorf("applyMigrate() sol = %v, want %v", tt.sol, tt.want)
			}
		})
	}
}

func TestPopulation_calcCosts(t *testing.T) {
	tests := []struct {
		name       string
		population Population
		want       Population
	}{
		{
			name:       "empty population",
			population: Population{},
			want:       Population{},
		},
		{
			name: "singleton population",
			population: Population{{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   0,
			}},
			want: Population{{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   10,
			}},
		},
		{
			name: "multi population",
			population: Population{
				{[]Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}}, 0},
				{[]Route{{Pos{0, 0, 0}, []Pos{{1, 0, 0}}}}, 0},
			},
			want: Population{
				{[]Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}}, 10},
				{[]Route{{Pos{0, 0, 0}, []Pos{{1, 0, 0}}}}, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.population.calcCosts()
			if !reflect.DeepEqual(tt.population, tt.want) {
				t.Errorf("calcCosts(): got %v, want %v", tt.population, tt.want)
			}
		})
	}
}

func TestSolution_calcCost(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want Solution
	}{
		{
			name: "0 routes",
			sol: Solution{
				Routes: []Route{},
				Cost:   0,
			},
			want: Solution{
				Routes: []Route{},
				Cost:   0,
			},
		},
		{
			name: "1 empty route",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{}}},
				Cost:   0,
			},
			want: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{}}},
				Cost:   0,
			},
		},
		{
			name: "1 singleton route",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   0,
			},
			want: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   10,
			},
		},
		{
			name: "1 route 3 activities",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
				Cost:   0,
			},
			want: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
				Cost:   4,
			},
		},
		{
			name: "2 singleton routes",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}},
					{Pos{0, 0, 0}, []Pos{{-3, -4, 0}}}},
				Cost: 0,
			},
			want: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}},
					{Pos{0, 0, 0}, []Pos{{-3, -4, 0}}}},
				Cost: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sol.calcCost()
			if !reflect.DeepEqual(tt.sol, tt.want) {
				t.Errorf("calcCost(): got %v, want %v", tt.sol, tt.want)
			}
		})
	}
}

func TestPopulation_getBest(t *testing.T) {
	tests := []struct {
		name         string
		population   Population
		wantSolution Solution
		wantErr      bool
	}{
		{
			name:         "empty population",
			population:   Population{},
			wantSolution: Solution{},
			wantErr:      true,
		},
		{
			name: "singleton population",
			population: Population{{
				Routes: []Route{},
				Cost:   1,
			}},
			wantSolution: Solution{
				Routes: []Route{},
				Cost:   1,
			},
			wantErr: false,
		},
		{
			name: "multiple same cost",
			population: Population{{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   1,
			}, {
				Routes: []Route{{Pos{5, 5, 0}, []Pos{{1, 1, 0}}}},
				Cost:   1,
			}},
			wantSolution: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   1,
			},
			wantErr: false,
		},
		{
			name: "multiple different cost",
			population: Population{{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   10,
			}, {
				Routes: []Route{{Pos{5, 5, 0}, []Pos{{1, 1, 0}}}},
				Cost:   1,
			}},
			wantSolution: Solution{
				Routes: []Route{{Pos{5, 5, 0}, []Pos{{1, 1, 0}}}},
				Cost:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSolution, err := tt.population.getBest()
			if (err != nil) != tt.wantErr {
				t.Errorf("getBest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSolution, tt.wantSolution) {
				t.Errorf("getBest() gotSolution = %v, want %v", gotSolution, tt.wantSolution)
			}
		})
	}
}

func TestPopulation_tournamentSelection(t *testing.T) {
	tests := []struct {
		name           string
		population     Population
		tournamentSize int
		want           Solution
		wantErr        bool
	}{
		{
			name:           "empty population",
			population:     Population{},
			tournamentSize: 5,
			want:           Solution{},
			wantErr:        true,
		},
		{
			name: "0 tournament size",
			population: Population{{
				Routes: []Route{},
				Cost:   1,
			}},
			tournamentSize: 0,
			want:           Solution{},
			wantErr:        true,
		},
		{
			name: "singleton population",
			population: Population{{
				Routes: []Route{},
				Cost:   1,
			}},
			tournamentSize: 1,
			want: Solution{
				Routes: []Route{},
				Cost:   1,
			},
			wantErr: false,
		},
		{
			name: "huge tournament size",
			population: Population{{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
				Cost:   10,
			}, {
				Routes: []Route{{Pos{5, 5, 0}, []Pos{{1, 1, 0}}}},
				Cost:   1,
			}},
			tournamentSize: 100,
			want: Solution{
				Routes: []Route{{Pos{5, 5, 0}, []Pos{{1, 1, 0}}}},
				Cost:   1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.population.tournamentSelection(tt.tournamentSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("tournamentSelection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tournamentSelection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_mutate_repeat(t *testing.T) {
	type testcase struct {
		inst         MDVRPInstance
		name         string
		sol          Solution
		maxMutations int
		mutationRate float64
		wantInst     MDVRPInstance
	}
	inst := MDVRPInstance{
		Activities: []Pos{
			{3, 5, 42},
			{6, 5, 42},
			{3, 4, 42},
			{10, 10, 69},
			{1, 2, 42},
			{33, 7, 69},
		},
		Depots:  []Pos{{5, 5, 42}, {9, 9, 50}},
		NRoutes: 2,
	}
	tests := make([]testcase, 100)
	for i := 0; i < 100; i++ {
		tests[i] = testcase{
			inst:         inst,
			name:         fmt.Sprintf("%d", i),
			sol:          randomSolution(inst),
			maxMutations: rand.Intn(5) + 1,
			mutationRate: rand.Float64(),
			wantInst:     inst,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sol.mutate(tt.inst, tt.maxMutations, tt.mutationRate)

			if tt.wantInst.NRoutes != len(tt.sol.Routes) {
				t.Errorf("mutate() gotRoutes = %d, wantNRoutes %d", len(tt.sol.Routes), tt.wantInst.NRoutes)
				return
			}
			if nAct := tt.sol.activityCount(); len(tt.wantInst.Activities) != nAct {
				t.Errorf("mutate() len(activities) = %d, len(inst.activities) %d", nAct, len(tt.wantInst.Activities))
				return
			}
			for _, route := range tt.sol.Routes {
				for _, activity := range route.Activities {
					if !fp.Contains(tt.wantInst.Activities, activity) {
						t.Errorf("mutate() got = %v, wantActivities %v", tt.sol, tt.wantInst.Activities)
						return
					}
				}
			}
			for _, activity := range tt.wantInst.Activities {
				has := false
				for _, route := range tt.sol.Routes {
					if fp.Contains(route.Activities, activity) {
						has = true
						break
					}
				}
				if !has {
					t.Errorf("mutate() got = %v, wantActivities %v", tt.sol, tt.wantInst.Activities)
					return
				}
			}
		})
	}
}

func TestSolution_removePoints(t *testing.T) {
	tests := []struct {
		name   string
		sol    Solution
		points []Pos
		want   Solution
	}{
		{
			name:   "empty Solution, empty points",
			sol:    Solution{},
			points: []Pos{},
			want:   Solution{},
		},
		{
			name:   "empty Solution, non-empty points",
			sol:    Solution{},
			points: []Pos{{0, 0, 0}},
			want:   Solution{},
		},
		{
			name:   "nonempty Solution, empty points",
			sol:    Solution{[]Route{{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}}, 0},
			points: []Pos{},
			want:   Solution{[]Route{{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}}, 0},
		},
		{
			name:   "point not in Solution",
			sol:    Solution{[]Route{{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}}, 0},
			points: []Pos{{42, 42, 0}},
			want:   Solution{[]Route{{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}}, 0},
		},
		{
			name:   "point in Solution",
			sol:    Solution{[]Route{{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}}, 0},
			points: []Pos{{1, 1, 0}},
			want:   Solution{[]Route{{Pos{0, 0, 0}, []Pos{}}}, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sol.removePoints(tt.points)
			if !reflect.DeepEqual(tt.sol, tt.want) {
				t.Errorf("removePoints() got = %v, want %v", tt.sol, tt.want)
			}
		})
	}
}