package genetic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			want: 0,
		},
		{
			name: "1 singleton route",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}}},
			},
			want: 11,
		},
		{
			name: "1 route, 3 activities, same zips",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
			},
			want: 4 + 0.4,
		},
		{
			name: "2 empty routes, diff depot",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{}},
					{Pos{0, 1, 1}, []Pos{}}},
				Cost: 0,
			},
			want: 0,
		},
		{
			name: "2 singleton routes, diff depot",
			sol: Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}},
					{Pos{0, 0, 1}, []Pos{{3, 4, 0}}}},
				Cost: 0,
			},
			want: 20 + 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sol.calcCost()
			assert.Equal(t, tt.want, tt.sol.Cost)
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
