package main

import (
	kMeans "bzone/backend/cmd/k-means"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	voronoi "github.com/pzsz/voronoi"
)

//	func VoronoiDiagram(clusters kMeans.Clusters) *voronoi.Diagram {
//		//convert clusters to vertices (input used by voronoi)
//		vertices := clusterToVertexList(clusters)
//		//Generate the voronoi diagram
//		diagram := generateVoronoiDiagram(vertices)
//		return diagram
//	}

var coordinateMultiplier = 10.0
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
		sites = append(sites, voronoi.Vertex{X: x * coordinateMultiplier, Y: y * coordinateMultiplier})
	}

	return sites
}

// func generateVoronoiDiagram(vertices []voronoi.Vertex) *voronoi.Diagram {
// 	bbox := voronoi.NewBBox(3.23, 7.3, 6.5, 5.92)
// 	voronoiDiagram := voronoi.ComputeDiagram(vertices, bbox, true)
// 	return voronoiDiagram
// }

func voronoiDiagram(clusters kMeans.Clusters) {
	// Set up the Voronoibox for the Netherlands
	bbox := voronoi.NewBBox(xl*coordinateMultiplier, xr*coordinateMultiplier, yt*coordinateMultiplier, yb*coordinateMultiplier)
	//sites := generateRandomPoints(50, bbox)

	//convert clusters to vertices
	sites := clusterToVertexList(clusters)

	v := voronoi.ComputeDiagram(sites, bbox, true)

	// Create the output image
	img := image.NewRGBA(image.Rect(0, 0, int(bbox.Xl-bbox.Xr), int(bbox.Yt-bbox.Yb)))
	drawSites(img, sites, bbox)
	drawEdges(img, v.Edges, bbox)

	// Save the image to a file
	f, err := os.Create("voronoi.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err)
	}
}

func generateRandomPoints(n int, bbox voronoi.BBox) []voronoi.Vertex {
	var points []voronoi.Vertex
	for i := 0; i < n; i++ {
		x := rand.Float64()*(bbox.Xl-bbox.Xr) + bbox.Xr
		fmt.Println(x)
		y := rand.Float64()*(bbox.Yt-bbox.Yb) + bbox.Yb
		points = append(points, voronoi.Vertex{x, y})
	}
	return points
}

func drawSites(img *image.RGBA, sites []voronoi.Vertex, bbox voronoi.BBox) {
	for _, site := range sites {
		x, y := int(site.X-bbox.Xr), int(site.Y-bbox.Yb)
		img.Set(x, y, color.Black)
	}
}

func drawEdges(img *image.RGBA, edges []*voronoi.Edge, bbox voronoi.BBox) {
	for _, edge := range edges {
		x1, y1 := int(edge.Va.X-bbox.Xr), int(edge.Va.Y-bbox.Yb)
		x2, y2 := int(edge.Vb.X-bbox.Xr), int(edge.Vb.Y-bbox.Yb)
		drawLine(img, x1, y1, x2, y2, color.Black)
	}
}

func drawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	dx := x2 - x1
	dy := y2 - y1
	if dx == 0 && dy == 0 {
		img.Set(x1, y1, c)
		return
	}
	if abs(dx) >= abs(dy) {
		if x2 < x1 {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		for x := x1; x <= x2; x++ {
			y := y1 + dy*(x-x1)/dx
			img.Set(x, y, c)
		}
	} else {
		if y2 < y1 {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		for y := y1; y <= y2; y++ {
			x := x1 + dx*(y-y1)/dy
			img.Set(x, y, c)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
