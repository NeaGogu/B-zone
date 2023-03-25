package genetic

import "math"

// Pos represents the position of an activity, with latitude and longitude coordinates and its Zipcode.
type Pos struct {
	Latitude  float64
	Longitude float64
	Zipcode   int
}

// dist calculates the Euclidean distance between Pos p0 and p1 based on Latitude and Longitude.
func dist(p0, p1 Pos) float64 {
	dx := p0.Latitude - p1.Latitude
	dy := p0.Longitude - p1.Longitude
	return math.Sqrt(dx*dx + dy*dy)
}

// copy copies the Latitude, Longitude, and Zipcode of a Pos and returns the copied Pos.
func (p *Pos) copy() Pos {
	return Pos{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		Zipcode:   p.Zipcode,
	}
}
