package genetic

import "math"

// Pos is the position of an activity
type Pos struct {
	Latitude  float64
	Longitude float64
	Zipcode   int
}

// dist calculates the Euclidean distance between Pos p0 and p1
func dist(p0, p1 Pos) float64 {
	return math.Sqrt(math.Pow(p1.Longitude-p0.Longitude, 2) + math.Pow(p1.Latitude-p0.Latitude, 2))
}

func (p *Pos) copy() Pos {
	return Pos{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		Zipcode:   p.Zipcode,
	}
}
