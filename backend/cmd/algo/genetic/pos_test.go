package genetic

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

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
			got := dist(tt.args.p0, tt.args.p1)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPos_copy(t *testing.T) {
	tests := []struct {
		name string
		pos  Pos
	}{
		{"blank Pos", Pos{}},
		{"all zero", Pos{0, 0, 0}},
		{"funny numbers", Pos{42, 69, 5318008}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.pos.copy(), tt.pos)
		})
	}
}
