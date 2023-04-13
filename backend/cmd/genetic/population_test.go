package genetic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantSolution, gotSolution)
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
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPopulation_calcCosts(t *testing.T) {
	tests := []struct {
		name       string
		population Population
		want       []float64
	}{
		{
			name:       "empty population",
			population: Population{},
			want:       []float64{},
		},
		{
			name: "singleton population",
			population: Population{Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}},
					{Pos{0, 0, 1}, []Pos{{3, 4, 0}}}},
			}},
			want: []float64{21},
		},
		{
			name: "double population",
			population: Population{Solution{
				Routes: []Route{{Pos{0, 0, 0}, []Pos{{3, 4, 0}}},
					{Pos{0, 0, 1}, []Pos{{3, 4, 0}}}},
			},
				Solution{
					Routes: []Route{{Pos{0, 0, 0}, []Pos{{0, 1, 0}, {1, 1, 0}, {1, 0, 0}}}},
				}},
			want: []float64{21, 4.4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.population.calcCosts()
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], tt.population[i].Cost)
			}
		})
	}
}
