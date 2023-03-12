package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/biogo/cluster/kmeans"
	//"github.com/biogo/cluster/kmeans"
	// "github.com/biogo/cluster/kmeans"
)

func main() {
	// I also uploaded these datasets so you can do further testing with them is you want
	files := []string{"Burma", "USCapitals", "Beergardens", "Cities", "Large"}
	for _, file := range files {
		fmt.Println(file)
		points := loadData("data\\" + file + ".txt")
		clusterPoints(points, 5, file, 250)
	}
	// data := loadCSVData("data/klant1_locs.csv")
	// clusterPoints(data, 14, "bumbal", 250)
}

func clusterPoints(points []*Pos, nClusters int, fileName string, maxDim int) {
	// This is for mapping from data coordinate space to image coordinate space
	// Not relevant for K-means
	minX, maxX := math.MaxFloat64, 0.0
	minY, maxY := math.MaxFloat64, 0.0
	for _, pos := range points {
		if pos.X > maxX {
			maxX = pos.X
		} else if pos.X < minX {
			minX = pos.X
		}
		if pos.Y > maxY {
			maxY = pos.Y
		} else if pos.Y < minY {
			minY = pos.Y
		}
	}
	var width, height int
	if xDiff, yDiff := maxX-minX, maxY-minY; xDiff > yDiff {
		height = maxDim
		width = int(float64(maxDim) * xDiff / yDiff)
	} else {
		width = maxDim
		height = int(float64(maxDim) * yDiff / xDiff)
	}

	// This is the important part for K-means
	// init K-means with the points
	km, err := kmeans.New(Points(points))
	if err != nil {
		panic(err)
	}
	// Seed K-means with nClusters center points
	// I think the center points can be manually initialized using km.SetCenters(c []cluster.Center)
	km.Seed(nClusters)
	// km.Cluster() runs the K-means clustering algorithm
	if err := km.Cluster(); err != nil {
		panic(err)
	}

	// The rest is for drawing the points on the image
	// some functions may be useful for seeing which points are in which cluster
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// loop over every cluster
	for _, c := range km.Centers() {
		col := randColor(r)
		// loop over every point in the cluster
		for _, i := range c.Members() {
			point := points[i]
			x := int((point.X - minX) * float64(width) / (maxX - minX))
			y := maxDim - int((point.Y-minY)*float64(height)/(maxY-minY))
			drawPoint(img, x, y, col)
		}
	}

	// save the image
	f, err := os.Create(fileName + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}

// loadCSVData loads the coordinate data from the csv data file
func loadCSVData(fileName string) []*Pos {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	var positions []*Pos
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		x, err0 := strconv.ParseFloat(nums[0], 64)
		y, err1 := strconv.ParseFloat(nums[1], 64)
		if err0 != nil {
			log.Fatal(err0)
		} else if err1 != nil {
			log.Fatal(err1)
		}
		positions = append(positions, &Pos{x, y})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return positions
}

// loadData loads the coordinate data from the txt data file
func loadData(fileName string) []*Pos {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "\\backend\\cmd\\k-means\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	scanner.Scan()
	var positions []*Pos
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)[1:]
		x, err0 := strconv.ParseFloat(nums[0], 64)
		y, err1 := strconv.ParseFloat(nums[1], 64)
		if err0 != nil {
			log.Fatal(err0)
		} else if err1 != nil {
			log.Fatal(err1)
		}
		positions = append(positions, &Pos{x, y})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return positions
}

type Pos struct {
	X float64
	Y float64
}

// Points implements cluster.Interface
type Points []*Pos

func (f Points) Len() int {
	return len(f)
}

// Values returns the values of f at index i as a []float64
func (f Points) Values(i int) []float64 {
	// I think if we want K-means to run on more parameters, we need to add it to the Pos struct and to the array here
	return []float64{f[i].X, f[i].Y}
}

// drawPoint draws a point on img. Not relevant to K-means
func drawPoint(img *image.RGBA, x int, y int, col color.Color) {
	img.Set(x, y, col)
}

// everything after here is for generating random numbers, which is not relevant to K-means

type IntRange struct {
	min, max int
}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func randColor(r *rand.Rand) color.RGBA {
	irc := IntRange{0, 256}
	return color.RGBA{R: uint8(irc.NextRandom(r)),
		G: uint8(irc.NextRandom(r)),
		B: uint8(irc.NextRandom(r)),
		A: 255}
}
