package haversine

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRadiusKm = 6371 // radius of the earth in kilometers.
	earthRadiusNM = 3440 // radius of the earth in nautical miles
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lon float64
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// nmToKm converts from nautical miles to kilometers.
func nmToKm(d float64) float64 {
	return d * earthRadiusKm / earthRadiusNM
}

// nmToMi converts from nautical miles to statute miles.
func nmToMi(d float64) float64 {
	return d * earthRadiusMi / earthRadiusNM
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in nautical miles.
func Distance(p, q Coord) (nm float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRadiusNM
}

// DistanceKm calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in kilometers.
func DistanceKm(p, q Coord) (km float64) {
	return nmToKm(Distance(p, q))
}

// DistanceMi calculates the shortest path between two coordinates on the surface
// of the Earth and returns the result in statute miles.
func DistanceMi(p, q Coord) (mi float64) {
	return nmToMi(Distance(p, q))
}
