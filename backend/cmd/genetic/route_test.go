package genetic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			got := greedyRoute(tt.args.depot, tt.args.points)
			assert.Equal(t, tt.want, got)
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
		{"i too high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 1, 0,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative i", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, -1, 0,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"j too high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative j", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"swap 2", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}, 0, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.applySwap(tt.i, tt.j)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.route)
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
		{"i to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 1, 0,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative i", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, -1, 0,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"j to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative j", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 0, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"2opt 2", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}}}, 0, 1,
			false, Route{Pos{0, 0, 0}, []Pos{{2, 2, 0}, {1, 1, 0}}}},
		{"2opt many", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}, {4, 4, 0}, {5, 5, 0}}}, 1, 3,
			false, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {4, 4, 0}, {3, 3, 0}, {2, 2, 0}, {5, 5, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.apply2Opt(tt.i, tt.j)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.route)
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
		{"i to high", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, 1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
		{"negative i", Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}, -1,
			true, Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.applyGreedy(tt.i)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.route)
			}
		})
	}
}

func TestRoute_applyInsertGreedy(t *testing.T) {
	tests := []struct {
		name  string
		route Route
		pos   Pos
		want  Route
	}{
		{
			name:  "empty route",
			route: Route{Pos{0, 0, 0}, []Pos{}},
			pos:   Pos{1, 1, 1},
			want:  Route{Pos{0, 0, 0}, []Pos{{1, 1, 1}}},
		},
		{
			name:  "insert before",
			route: Route{Pos{0, 0, 0}, []Pos{{1, 1, 0}, {0, 1, 0}}},
			pos:   Pos{1, 0, 0},
			want:  Route{Pos{0, 0, 0}, []Pos{{1, 0, 0}, {1, 1, 0}, {0, 1, 0}}},
		},
		{
			name:  "insert middle",
			route: Route{Pos{0, 0, 0}, []Pos{{1, 0, 0}, {0, 1, 0}}},
			pos:   Pos{1, 1, 0},
			want:  Route{Pos{0, 0, 0}, []Pos{{1, 0, 0}, {1, 1, 0}, {0, 1, 0}}},
		},
		{
			name:  "insert after",
			route: Route{Pos{0, 0, 0}, []Pos{{1, 0, 0}, {1, 1, 0}}},
			pos:   Pos{0, 1, 0},
			want:  Route{Pos{0, 0, 0}, []Pos{{1, 0, 0}, {1, 1, 0}, {0, 1, 0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.route.applyInsertGreedy(tt.pos)
			assert.Equal(t, tt.want, tt.route)
		})
	}
}

func TestRoute_applyChangeDepot(t *testing.T) {
	type args struct {
		inst MDMTSPInstance
		d    int
	}
	tests := []struct {
		name    string
		route   Route
		args    args
		wantErr bool
		want    Route
	}{
		{"d too large", Route{Pos{1, 1, 1}, []Pos{}}, args{MDMTSPInstance{Depots: []Pos{}}, 1}, true, Route{}},
		{"negative d", Route{Pos{1, 1, 1}, []Pos{}}, args{MDMTSPInstance{Depots: []Pos{}}, -1}, true, Route{}},
		{
			name:    "single depot",
			route:   Route{Pos{1, 1, 1}, []Pos{}},
			args:    args{MDMTSPInstance{Depots: []Pos{{1, 1, 1}}}, 0},
			wantErr: false,
			want:    Route{Pos{1, 1, 1}, []Pos{}},
		},
		{
			name:    "multiple depots change",
			route:   Route{Pos{1, 1, 1}, []Pos{}},
			args:    args{MDMTSPInstance{Depots: []Pos{{1, 1, 1}, {2, 2, 2}}}, 1},
			wantErr: false,
			want:    Route{Pos{2, 2, 2}, []Pos{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.route.applyChangeDepot(tt.args.inst, tt.args.d)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, tt.route)
			}
		})
	}
}

func TestRoute_copy(t *testing.T) {
	tests := []struct {
		name  string
		route Route
	}{
		{"empty route", Route{Pos{1, 1, 1}, []Pos{}}},
		{"singleton route", Route{Pos{1, 1, 1}, []Pos{{2, 2, 2}}}},
		{"large route", Route{Pos{1, 1, 1}, []Pos{{2, 2, 2}, {3, 3, 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.route, tt.route.copy())
		})
	}
}
