package main

import (
	kMeans "bzone/backend/cmd/k-means"
	"fmt"
	"os"

	s2 "github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	voronoi "github.com/pzsz/voronoi"
)

//	func VoronoiDiagram(clusters kMeans.Clusters) *voronoi.Diagram {
//		//convert clusters to vertices (input used by voronoi)
//		vertices := clusterToVertexList(clusters)
//		//Generate the voronoi diagram
//		diagram := generateVoronoiDiagram(vertices)
//		return diagram
//	}

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
		point := s2.PointFromLatLng(s2.LatLngFromDegrees(x, y))
		sites = append(sites, voronoi.Vertex{X: point.X, Y: point.Y})
	}

	return sites
}

func voronoiDiagram(clusters kMeans.Clusters) (*geojson.FeatureCollection, error) {
	// Set up the Voronoibox for the Netherlands
	//bbox := voronoi.NewBBox(xl*coordinateMultiplier, xr*coordinateMultiplier, yt*coordinateMultiplier, yb*coordinateMultiplier)
	// Generate the Voronoi diagram using the voronoi library
	bbox := voronoi.NewBBox(-180.0, 180.0, -90.0, 90.0)

	//sites := generateRandomPoints(50, bbox)

	//convert clusters to vertices
	sites := clusterToVertexList(clusters)

	voronoi := voronoi.ComputeDiagram(sites, bbox, true)

	featureCollectionVoronoi := voronoiDiagramToFeature(voronoi)

	jsonResult, err := featureCollectionToJSON(featureCollectionVoronoi)
	if err != nil {
		return nil, fmt.Errorf("Error in featureColletionToJSON: %v", err)
	}
	fmt.Println(string(jsonResult))
	return featureCollectionVoronoi, nil
}

func voronoiDiagramToFeature(vd *voronoi.Diagram) *geojson.FeatureCollection {
	// Convert the Voronoi diagram to a GeoJSON feature collection
	fc := geojson.NewFeatureCollection()
	for _, edge := range vd.Edges {
		// fmt.Println(fmt.Sprintf("the edge leftcell is: %v", edge.LeftCell))
		// fmt.Println(fmt.Sprintf("the edge rightcell is: %v", edge.RightCell))
		// fmt.Println(fmt.Sprintf("the edge Va (startvertex) is: %v", edge.Va))
		// fmt.Println(fmt.Sprintf("the edge Vb (endvertex) is: %v", edge.Vb))

		line := geojson.NewLineStringFeature([][]float64{featureFromVertex(edge.Va), featureFromVertex(edge.Vb)})
		fmt.Println(fc)
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
	// fmt.Println(fmt.Sprintf("the vertex X: %v", v.Vertex.X))
	// fmt.Println(fmt.Sprintf("the vertex Y: %v", v.Vertex.Y))
	// fmt.Println(fmt.Sprintf("latLng: %v", s2.LatLngFromDegrees(v.Vertex.X, v.Vertex.Y)))
	return []float64{v.Vertex.X, v.Vertex.Y}
}
