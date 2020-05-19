package haversine

import (
	"math"
)

const (
	EarthRadiusMi = 3958 // radius of the earth in miles.
	EarthRadiusKm = 6371 // radius of the earth in kilometers.
	EarthRadiusNM = 3440 // radius of the earth in nautical miles
)

// Coord represents a lat/long geographic coordinate, usually in degrees +EN/-WS.
type Coord struct {
	Lat float64
	Lon float64
}

// Radians converts a radians-based Coord to a degrees-based Coord 
func (c *Coord) Degrees() Coord {
	return Coord{c.Lat * 180 / math.Pi, c.Lon * 180 / math.Pi}
}

// Radians converts a degrees-based Coord to a radians-based Coord 
func (c *Coord) Radians() Coord {
	return Coord{c.Lat * math.Pi / 180, c.Lon * math.Pi / 180}
}

// NmToMi converts from nautical miles to statute miles.
func NmToMi(d float64) float64 {
	return d * EarthRadiusMi / EarthRadiusNM
}

// NmToKm converts from nautical miles to kilometers.
func NmToKm(d float64) float64 {
	return d * EarthRadiusKm / EarthRadiusNM
}

// KmToNm converts from kilometers to nautical miles.
func KmToNm(d float64) float64 {
	return d * EarthRadiusNM / EarthRadiusKm
}

// IntAgle calculates the internal angle between two coordinates on a surface
// and returns the result in radians
func IntAngle(p, q Coord) float64 {
	point1 := p.Radians()
	point2 := q.Radians()

	diffLat := point2.Lat - point1.Lat
	diffLon := point2.Lon - point1.Lon

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(point1.Lat)*math.Cos(point2.Lat)*
		math.Pow(math.Sin(diffLon/2), 2)

	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// Distance calculates the shortest (aka great circle) arc between two coordinates on a sphere
// of a given radius and returns the resulting great circle arc length
func Distance(p, q Coord, radius float64) (gc float64) {

	c := IntAngle(p, q)

	return c * radius
}

// DistanceMi calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in statute miles.
func DistanceMi(p, q Coord) (mi float64) {
	return Distance(p, q, EarthRadiusMi)
}

// DistanceKm calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in kilometers.
func DistanceKm(p, q Coord) (km float64) {
	return Distance(p, q, EarthRadiusKm)
}

// DistanceNM calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in nautical miles.
func DistanceNM(p, q Coord) (nm float64) {
	return Distance(p, q, EarthRadiusNM)
}
