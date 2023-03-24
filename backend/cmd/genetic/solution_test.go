package genetic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
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
			err := tt.sol.applySwap(tt.args.r0, tt.args.i, tt.args.r1, tt.args.j)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.sol)
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
			err := tt.sol.applyMigrate(tt.args.r0, tt.args.i, tt.args.r1, tt.args.j)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.sol)
			}
		})
	}
}

func TestSolution_calcCost(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want float64
	}{
		{
			name: "0 routes",
			sol: Solution{
				Routes: []Route{},
			},
			want: 0,
		},
		{
			name: "1 empty route",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{}}},
			},
			want: 100,
		},
		{
			name: "1 singleton route",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
			},
			want: 10 + 10 + 100,
		},
		{
			name: "1 route, 3 activities, same zips",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
			},
			want: 4 + 10 + 100,
		},
		{
			name: "1 route, 3 activities, diff zips",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 1}, {1, 0, 2}}}},
			},
			want: 4 + 30 + 100,
		},
		{
			name: "2 empty routes, same depot",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{}},
					{Pos{0, 1, 1}, []Pos{}}},
				Cost: 0,
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sol.calcCost()
			assert.Equal(t, tt.want, tt.sol.Cost)
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
			if !assert.Equal(t, tt.wantInst.NRoutes, len(tt.sol.Routes)) {
				return
			}
			var gotActivities []Pos
			for _, route := range tt.sol.Routes {
				for _, activity := range route.Activities {
					gotActivities = append(gotActivities, activity)
				}
			}
			assert.ElementsMatch(t, tt.inst.Activities, gotActivities)
		})
	}
}

func TestSolution_removeActivities(t *testing.T) {
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
			assert.Equal(t, tt.want, tt.sol)
		})
	}
}

func TestSolution_copy(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
	}{
		{"0 routes", Solution{[]Route{}, 42}},
		{"1 routes", Solution{[]Route{{Pos{42, 69, 420}, []Pos{{1, 2, 3}}}}, 42}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.sol, tt.sol.copy())
		})
	}
}

func TestSolution_zipsPerRouteSum(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want int
	}{
		{"no routes", Solution{[]Route{}, 42}, 0},
		{"1 singleton route", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}}, 42}, 1},
		{"1 route, 2 activities same zip", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}, {1, 0, 42}}}}, 42}, 1},
		{"1 route, 2 activities diff zip", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}, {1, 0, 69}}}}, 42}, 2},
		{"2 routes, 1 activity per route same zip", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}, {Pos{0, 0, 42}, []Pos{{1, 0, 42}}}}, 42}, 2},
		{"2 routes, 1 activity per route diff zip", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}, {Pos{0, 0, 42}, []Pos{{1, 0, 69}}}}, 42}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sol.zipsPerRouteSum())
		})
	}
}

func TestSolution_nDepots(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want int
	}{
		{"no routes", Solution{[]Route{}, 42}, 0},
		{"1 singleton route", Solution{[]Route{{Pos{0, 0, 42}, []Pos{}}}, 42}, 1},
		{"2 routes, same depot", Solution{[]Route{{Pos{0, 0, 42}, []Pos{}}, {Pos{0, 0, 42}, []Pos{}}}, 42}, 1},
		{"2 routes, diff depot", Solution{[]Route{{Pos{0, 0, 42}, []Pos{}}, {Pos{0, 0, 69}, []Pos{}}}, 42}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sol.nDepots())
		})
	}
}

func TestSolution_routeLengthVariance(t *testing.T) {
	tests := []struct {
		name string
		sol  Solution
		want float64
	}{
		{"no routes", Solution{[]Route{}, 42}, 0},
		{"1 route", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}}, 42}, 0},
		{"2 routes, equal length", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}, {Pos{0, 0, 42}, []Pos{{0, 0, 42}}}}, 42}, 0},
		{"2 routes, diff length", Solution{[]Route{{Pos{0, 0, 42}, []Pos{{0, 0, 42}}}, {Pos{0, 0, 42}, []Pos{{0, 0, 42}, {0, 0, 42}}}}, 42}, 0.5},
		{"3 routes, diff length", Solution{[]Route{{Pos{0, 0, 42}, []Pos{}}, {Pos{0, 0, 42}, []Pos{}}, {Pos{0, 0, 42}, []Pos{{0, 0, 42}, {0, 0, 42}, {0, 0, 42}}}}, 42}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.sol.routeLengthVariance())
		})
	}
}

func Test_randomSolution_repeat(t *testing.T) {
	type testcase struct {
		inst     MDMTSPInstance
		name     string
		wantInst MDMTSPInstance
	}
	tests := make([]testcase, 100)
	for i := 0; i < 100; i++ {
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
			NRoutes: rand.Intn(5) + 2,
		}
		tests[i] = testcase{
			inst:     inst,
			name:     fmt.Sprintf("%d", i),
			wantInst: inst,
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sol := randomSolution(tt.inst)
			if !assert.Equal(t, tt.wantInst.NRoutes, len(sol.Routes)) {
				return
			}
			var gotActivities []Pos
			for _, route := range sol.Routes {
				for _, activity := range route.Activities {
					gotActivities = append(gotActivities, activity)
				}
			}
			assert.ElementsMatch(t, tt.wantInst.Activities, gotActivities)
		})
	}
}

func TestSolution_applyMigrateZip(t *testing.T) {
	type args struct {
		r0  int
		r1  int
		zip int
	}
	tests := []struct {
		name    string
		sol     Solution
		args    args
		wantErr bool
		want    Solution
	}{
		{"negative r0", Solution{Routes: []Route{}}, args{-1, 0, 0}, true, Solution{}},
		{"negative r1", Solution{Routes: []Route{}}, args{0, -1, 0}, true, Solution{}},
		{"r0 too large", Solution{Routes: []Route{}}, args{1, 0, 0}, true, Solution{}},
		{"r0 too large", Solution{Routes: []Route{}}, args{0, 1, 0}, true, Solution{}},
		{"r0==r1", Solution{Routes: []Route{}}, args{0, 0, 0}, true, Solution{}},
		{"singleton r0, empty r1", Solution{Routes: []Route{{Activities: []Pos{{0, 0, 0}}}, {Activities: []Pos{}}}}, args{0, 1, 0}, false,
			Solution{Routes: []Route{{}, {Activities: []Pos{{0, 0, 0}}}}}},
		{"r0 3 same zip, empty r1", Solution{Routes: []Route{{Activities: []Pos{{0, 1, 0}, {1, 1, 0}}}, {Activities: []Pos{}}}}, args{0, 1, 0}, false,
			Solution{Routes: []Route{{}, {Activities: []Pos{{1, 1, 0}, {0, 1, 0}}}}}},
		{"r0 2 diff zip, empty r1", Solution{Routes: []Route{{Activities: []Pos{{0, 0, 0}, {1, 1, 1}}}, {Activities: []Pos{}}}}, args{0, 1, 0}, false,
			Solution{Routes: []Route{{Activities: []Pos{{1, 1, 1}}}, {Activities: []Pos{{0, 0, 0}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sol.applyMigrateZip(tt.args.r0, tt.args.r1, tt.args.zip)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.sol)
			}
		})
	}
}
