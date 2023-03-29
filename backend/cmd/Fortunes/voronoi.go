package main

import (
	kMeans "bzone/backend/cmd/k-means"

	voronoi "github.com/pzsz/voronoi"
)

func VoronoiDiagram(clusters kMeans.Clusters) *voronoi.Diagram {
	//convert clusters to vertices (input used by voronoi)
	vertices := clusterToVertexList(clusters)
	//Generate the voronoi diagram
	diagram := generateVoronoiDiagram(vertices)
	return diagram
}

func clusterToVertexList(clusters kMeans.Clusters) []voronoi.Vertex {
	var sites []voronoi.Vertex

	for _, cluster := range clusters {
		sites = append(sites, voronoi.Vertex{X: cluster.Center.Latitude, Y: cluster.Center.Longitude})
	}
	return sites
}

func generateVoronoiDiagram(vertices []voronoi.Vertex) *voronoi.Diagram {
	bbox := voronoi.NewBBox(3.23, 7.3, 6.5, 5.92)
	voronoiDiagram := voronoi.ComputeDiagram(vertices, bbox, true)
	return voronoiDiagram
}

// func createVoronoiImage(width, height, numSites int) image.Image {
// 	diagram := generateVoronoiDiagram(width, height, numSites)

// 	img := image.NewRGBA(image.Rect(0, 0, width, height))
// 	for _, edge := range diagram.Edges {

// 		x0, y0 := int(edge.Va.X), int(edge.Va.Y)
// 		x1, y1 := int(edge.Vb.X), int(edge.Vb.Y)

// 		drawLine(img, x0, y0, x1, y1, color.Black)
// 	}

// 	return img
// }

// func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
// 	dx := abs(x1 - x0)
// 	dy := abs(y1 - y0)
// 	sx := -1
// 	if x0 < x1 {
// 		sx = 1
// 	}
// 	sy := -1
// 	if y0 < y1 {
// 		sy = 1
// 	}
// 	err := dx - dy
// 	for {
// 		img.Set(x0, y0, c)
// 		if x0 == x1 && y0 == y1 {
// 			break
// 		}
// 		e2 := 2 * err
// 		if e2 > -dy {
// 			err -= dy
// 			x0 += sx
// 		}
// 		if e2 < dx {
// 			err += dx
// 			y0 += sy
// 		}
// 	}
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func main() {

// 	width := 800
// 	height := 600
// 	numSites := 100

// 	img := createVoronoiImage(width, height, numSites)
// 	file, err := os.Create("voronoi.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	png.Encode(file, img)
// }
