package genetic

import (
	"reflect"
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
