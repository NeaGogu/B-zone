package genetic

import (
	"reflect"
	"testing"
)

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
