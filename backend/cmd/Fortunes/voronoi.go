package voronoi

import (
	kMeans "bzone/backend/cmd/k-means"
	"fmt"
	"os"

	geojson "github.com/paulmach/go.geojson"
	voronoi "github.com/pzsz/voronoi"
)

var coordinateMultiplier = 100.0
var xl = 3.1
var xr = 7.1
var yt = 50.6
var yb = 53.7

func clusterToVertexList(clusters kMeans.Clusters) []voronoi.Vertex {
	var sites []voronoi.Vertex
	for _, cluster := range clusters {
		//check if x coordinate is within the Netherlands
		x := cluster.Center.Longitude
		if x < xl {
			fmt.Fprintf(os.Stderr, "x coordinate too small, adjusting to the border: %f", x)
			x = xl
		}
		if x > xr {
			fmt.Fprintf(os.Stderr, "x coordinate too large, adjusting to the border: %f", x)
			x = xr
		}

		//check if y coordinate is within the Netherlands
		y := cluster.Center.Latitude
		if y < yt {
			fmt.Fprintf(os.Stderr, "y coordinate too small, adjusting to the border: %f", y)
			y = yt
		}
		if y > yb {
			fmt.Fprintf(os.Stderr, "y coordinate too large, adjusting to the border: %f", y)
			y = yb
		}

		sites = append(sites, voronoi.Vertex{X: x, Y: y})
	}

	return sites
}

func VoronoiDiagram(clusters kMeans.Clusters) (*geojson.FeatureCollection, error) {
	// Generate the Voronoi diagram using the voronoi library
	bbox := voronoi.NewBBox(-180.0, 180.0, -90.0, 90.0)

	//convert clusters to vertices
	sites := clusterToVertexList(clusters)

	voronoi := voronoi.ComputeDiagram(sites, bbox, true)

	featureCollectionVoronoi := voronoiDiagramToFeature(voronoi)

	return featureCollectionVoronoi, nil
}

func voronoiDiagramToFeature(vd *voronoi.Diagram) *geojson.FeatureCollection {
	// Convert the Voronoi diagram to a GeoJSON feature collection
	fc := geojson.NewFeatureCollection()
	for _, edge := range vd.Edges {
		line := geojson.NewLineStringFeature([][]float64{featureFromVertex(edge.Va), featureFromVertex(edge.Vb)})

		fc.AddFeature(line)
	}
	return fc
}

func featureCollectionToJSON(fc *geojson.FeatureCollection) ([]byte, error) {
	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("Error in featureColletionToJSON: %v", err)
	}
	return rawJSON, nil
}

func featureFromVertex(v voronoi.EdgeVertex) []float64 {
	return []float64{v.Vertex.X, v.Vertex.Y}
}
