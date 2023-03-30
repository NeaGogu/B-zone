package main

import (
	kMeans "bzone/backend/cmd/k-means"
	"fmt"
	"os"
	"syscall/js"

	leaflet "github.com/ctessum/go-leaflet"
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

func voronoiDiagram(clusters kMeans.Clusters) *voronoi.Diagram {
	// Set up the Voronoibox for the Netherlands
	//bbox := voronoi.NewBBox(xl*coordinateMultiplier, xr*coordinateMultiplier, yt*coordinateMultiplier, yb*coordinateMultiplier)
	// Generate the Voronoi diagram using the voronoi library
	bbox := voronoi.NewBBox(-180.0, 180.0, -90.0, 90.0)

	//sites := generateRandomPoints(50, bbox)

	//convert clusters to vertices
	sites := clusterToVertexList(clusters)

	v := voronoi.ComputeDiagram(sites, bbox, true)

	return v
}

func drawVoronoiDiagramOnMap(vd *voronoi.Diagram, mapElementID string) {
	// Convert the Voronoi diagram to a GeoJSON feature collection
	fc := geojson.NewFeatureCollection()
	for _, edge := range vd.Edges {
		line := geojson.NewLineStringFeature([][]float64{featureFromVertex(edge.Va), featureFromVertex(edge.Vb)})
		fc.AddFeature(line)
	}

	// Draw the Voronoi diagram on a Leaflet map
	mapElement := js.Global().Get("document").Call("getElementById", mapElementID)
	mapV := leaflet.NewMap(mapElement)
	layer := leaflet.NewGeoJSONLayer(fc, leaflet.GeoJSONLayerOptions{})
	mapV.AddLayer(layer)
}

func featureFromVertex(v voronoi.EdgeVertex) []float64 {
	return []float64{v.Vertex.X, v.Vertex.Y}
}
