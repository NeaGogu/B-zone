package genetic

import (
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math/rand"
	"reflect"
	"testing"
)

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

func TestSolution_mutate_repeat(t *testing.T) {
	type testcase struct {
		inst         MDMTSPInstance
		name         string
		sol          Solution
		maxMutations int
		mutationRate float64
		wantInst     MDMTSPInstance
	}
	inst := MDMTSPInstance{
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
			if nAct := tt.sol.activityCount(t); len(tt.wantInst.Activities) != nAct {
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
			tt.sol.removeActivities(tt.points)
			if !reflect.DeepEqual(tt.sol, tt.want) {
				t.Errorf("removeActivities() got = %v, want %v", tt.sol, tt.want)
			}
		})
	}
}

func (sol *Solution) activityCount(t testing.TB) int {
	t.Helper()
	return fp.Reduce(sol.Routes,
		func(route Route, count int) int {
			return count + len(route.Activities)
		}, 0)
}
