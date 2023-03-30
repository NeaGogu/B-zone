package main

import (
	"testing"

	kMeans "bzone/backend/cmd/k-means"

	voronoi "github.com/pzsz/voronoi"
	"github.com/stretchr/testify/assert"
)

func TestClusterToVertexList(t *testing.T) {
	// Set up test data
	tests := []struct {
		name     string
		clusters kMeans.Clusters
		want     []voronoi.Vertex
	}{
		{
			name: "Test with two clusters within bounds",
			clusters: kMeans.Clusters{
				{
					Center: kMeans.Coordinates{Latitude: 52.370216, Longitude: 4.895168},
				},
				{
					Center: kMeans.Coordinates{Latitude: 51.9194, Longitude: 4.4818},
				},
			},
			want: []voronoi.Vertex{
				{X: 48.9516800, Y: 523.702160},
				{X: 44.8180000, Y: 519.194000},
			},
		},
		{
			name: "Test with one cluster outside bounds",
			clusters: kMeans.Clusters{
				{
					Center: kMeans.Coordinates{Latitude: 53.4808, Longitude: -2.2426},
				},
			},
			want: []voronoi.Vertex{
				{X: 31, Y: 534.808},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clusterToVertexList(tt.clusters)
			// GIS/handheld gps precision
			epsilon := 0.0001
			// Check the result
			if len(tt.want) != len(got) {
				t.Fatalf("unexpected result length: got %d, want %d", len(got), len(tt.want))
			}

			for i, v := range got {
				// if v.X < xl*coordinateMultiplier || v.X > xr*coordinateMultiplier || v.Y < yb*coordinateMultiplier || v.Y > yt*coordinateMultiplier {
				// 	t.Errorf("result vertex %d is outside the Netherlands: %v", i, v)
				// }

				assert.InDelta(t, tt.want[i].X, v.X, epsilon, "X coordinates are not equal")
				assert.InDelta(t, tt.want[i].Y, v.Y, epsilon, "Y coordinates are not equal")
			}
		})
	}
}

func TestVoronoiDiagram(t *testing.T) {
	clusters := kMeans.Clusters{
		{Center: kMeans.Coordinates{53.171767, 6.532410}}, //groningen
		{Center: kMeans.Coordinates{52.313816, 4.906434}}, //amsterdam
		{Center: kMeans.Coordinates{52.233807, 6.951607}}, //enschede
		{Center: kMeans.Coordinates{51.668418, 5.298165}}, //den bosch
		{Center: kMeans.Coordinates{51.9202, 4.4818}},     //Rotterdam
		{Center: kMeans.Coordinates{52.0705, 4.3007}},     //den haag
		{Center: kMeans.Coordinates{52.0907, 5.1214}},     //utrecht
		{Center: kMeans.Coordinates{51.4416, 5.4697}},     //Eindhoven
		{Center: kMeans.Coordinates{51.5555, 5.0913}},     //tilburg
		{Center: kMeans.Coordinates{51.5831, 4.7766}},     //Breda
		{Center: kMeans.Coordinates{51.8126, 5.8372}},     //nimma
	}
	result, err := voronoiDiagram(clusters)

	if err != nil {
		t.Errorf("didn't expect error but got: %v", err)
	}
	t.Errorf("result is : %v", result)
}
