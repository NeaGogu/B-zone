package main

import (
	"fmt"
	fp "github.com/rjNemo/underscore"
	"testing"
)

func Test_crossover_repeat(t *testing.T) {
	type args struct {
		parent1 Solution
		parent2 Solution
	}
	type testcase struct {
		name     string
		args     args
		wantInst VRPInstance
	}
	inst := VRPInstance{
		activities: []Pos{
			{3, 5, 42},
			{6, 5, 42},
			{3, 4, 42},
			{10, 10, 69},
			{1, 2, 42},
			{33, 7, 69},
		},
		depots:  []Pos{{5, 5, 42}, {9, 9, 50}},
		nRoutes: 2,
	}
	tests := make([]testcase, 100)
	for i := 0; i < 100; i++ {
		tests[i] = testcase{
			name: fmt.Sprintf("%d", i),
			args: args{
				parent1: randomSolution(inst),
				parent2: randomSolution(inst),
			},
			wantInst: inst,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := crossover(tt.args.parent1, tt.args.parent2)

			if tt.wantInst.nRoutes != len(got.routes) {
				t.Errorf("crossover() gotRoutes = %d, wantNRoutes %d", len(got.routes), tt.wantInst.nRoutes)
				return
			}
			if nAct := got.activityCount(); len(tt.wantInst.activities) != nAct {
				t.Errorf("crossover() len(activities) = %d, len(inst.activities) %d", nAct, len(tt.wantInst.activities))
				return
			}
			for _, route := range got.routes {
				for _, activity := range route.activities {
					if !fp.Contains(tt.wantInst.activities, activity) {
						t.Errorf("crossover() got = %v, wantActivities %v", got, tt.wantInst.activities)
						return
					}
				}
			}
			for _, activity := range tt.wantInst.activities {
				has := false
				for _, route := range got.routes {
					if fp.Contains(route.activities, activity) {
						has = true
						break
					}
				}
				if !has {
					t.Errorf("crossover() got = %v, wantActivities %v", got, tt.wantInst.activities)
					return
				}
			}
		})
	}
}
