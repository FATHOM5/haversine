package haversine_test

import (
	"math"
	"testing"

	"github.com/FATHOM5/haversine"
)

var tests = []struct {
	p     haversine.Coord
	q     haversine.Coord
	outMi float64
	outKm float64
}{
	{
		haversine.Coord{Lat: 22.55, Lon: 43.12},  // Rio de Janeiro, Brazil
		haversine.Coord{Lat: 13.45, Lon: 100.28}, // Bangkok, Thailand
		3786.2512,
		6094.5444,
	},
	{
		haversine.Coord{Lat: 20.10, Lon: 57.30}, // Port Louis, Mauritius
		haversine.Coord{Lat: 0.57, Lon: 100.21}, // Padang, Indonesia
		3196.6710,
		5145.5257,
	},
	{
		haversine.Coord{Lat: 51.45, Lon: 1.15},  // Oxford, United Kingdom
		haversine.Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		863.0311,
		1389.1793,
	},
	{
		haversine.Coord{Lat: 22.34, Lon: 17.05}, // Windhoek, Namibia
		haversine.Coord{Lat: 51.56, Lon: 4.29},  // Rotterdam, Netherlands
		2130.8298,
		3429.8931,
	},
	{
		haversine.Coord{Lat: 63.24, Lon: 56.59}, // Esperanza, Argentina
		haversine.Coord{Lat: 8.50, Lon: 13.14},  // Luanda, Angola
		4346.3983,
		6996.1859,
	},
	{
		haversine.Coord{Lat: 90.00, Lon: 0.00}, // North/South Poles
		haversine.Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		2866.1346,
		4613.4775,
	},
	{
		haversine.Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		haversine.Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		6261.0527,
		10078.1119,
		// 10070.0000,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		mi := haversine.DistanceMi(input.p, input.q)
		km := haversine.DistanceKm(input.p, input.q)
		// Note: Trunc and mult by 10000 eliminate floating point errors
		// past four decimal places.
		mi = math.Trunc(mi * 10000)
		km = math.Trunc(km * 10000)

		if input.outMi*10000 != mi || input.outKm*10000 != km {
			t.Errorf("fail: want %v %v -> %v %v got %f %f",
				input.p,
				input.q,
				input.outMi,
				input.outKm,
				mi,
				km,
			)
		}
	}
}
