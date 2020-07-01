package haversine_test

import (
	"math"
	"testing"

	"github.com/loraxipam/haversine"
)

// Great circle calculations are really only valid for about four sig figs
// for Earth, anyway. Up this, (or should that be, "decrease this"?)
// for real spheres.
const ppm = 0.0001

var tests = []struct {
	p     haversine.Coord
	q     haversine.Coord
	outMi float64
	outKm float64
	outRa float64
	town  string
}{
	{
		haversine.Coord{Lat: -22.55, Lon: -43.12}, // Rio de Janeiro, Brazil
		haversine.Coord{Lat: 13.45, Lon: 100.28},  // Bangkok, Thailand
		9956,
		16026,
		2.515463,
		"Rio to Bangkok",
	},
	{
		haversine.Coord{Lat: -20.10, Lon: 57.30}, // Port Louis, Mauritius
		haversine.Coord{Lat: 0.57, Lon: 100.21},  // Padang, Indonesia
		3233,
		5205,
		0.817067,
		"Mauritius to Indonesia",
	},
	{
		haversine.Coord{Lat: 51.45, Lon: -1.15}, // Oxford, United Kingdom
		haversine.Coord{Lat: 41.54, Lon: 12.27}, // Vatican, City Vatican City
		933,
		1501,
		0.235736,
		"Oxford to the Vatican",
	},
	{
		// These are antipodes
		haversine.Coord{Lat: 32.30, Lon: -64.77},  // Bermuda
		haversine.Coord{Lat: -32.30, Lon: 115.23}, // Perth, Australia
		12434,   // earthRadiusMi * pi
		20015,   // earthRadiusKm * pi
		3.14159, // pi
		"Bermuda to Perth",
	},
	{
		haversine.Coord{Lat: -22.34, Lon: 17.05}, // Windhoek, Namibia
		haversine.Coord{Lat: 51.56, Lon: 4.29},   // Rotterdam, Netherlands
		5163,
		8311,
		1.304548,
		"Namibia to the Netherlands",
	},
	{
		haversine.Coord{Lat: -63.24, Lon: -56.59}, // Esperanza, Argentina
		haversine.Coord{Lat: -8.50, Lon: 13.14},   // Luanda, Angola
		5068,
		8157,
		1.280482,
		"Argentina to Angola",
	},
	{
		haversine.Coord{Lat: 90.00, Lon: 0.00}, // North Pole
		haversine.Coord{Lat: 48.51, Lon: 2.21}, // Paris,  France
		2866,
		4613,
		0.724137,
		"Santa to Paris",
	},
	{
		haversine.Coord{Lat: -90.00, Lon: 0.00}, // South Pole
		haversine.Coord{Lat: 48.51, Lon: 2.21},  // Paris,  France
		9568,
		15401,
		2.417456,
		"Penguins to Paris",
	},
	{
		haversine.Coord{Lat: 45.04, Lon: 7.42},  // Turin, Italy
		haversine.Coord{Lat: 3.09, Lon: 101.42}, // Kuala Lumpur, Malaysia
		6261,
		10078,
		1.581873,
		"Turin to Malaysia",
	},
	{
		haversine.Coord{Lat: 45.71, Lon: -122.43}, // Hockinson, Washington, USA
		haversine.Coord{Lat: 29.13, Lon: -80.96},  // Wilbur-by-the-Sea, Florida, USA
		2510,
		4040,
		0.634270,
		"Washington to Florida",
	},
}

// TestHaversineDistance makes sure that the several great circle distance
// functions work properly
func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		mi, _ := math.Modf(haversine.DistanceMi(input.p, input.q))
		km, _ := math.Modf(haversine.DistanceKm(input.p, input.q))

		if input.outMi != mi || input.outKm != km {
			t.Errorf("fail: want %v %v -> %v %v got %g %g %s",
				input.p,
				input.q,
				input.outMi,
				input.outKm,
				mi,
				km,
				input.town,
			)
		}
	}
}

// TestIntAngle makes sure the internal angle calculation result is within tolerance
func TestIntAngle(t *testing.T) {
	for _, input := range tests {
		iAngle := haversine.IntAngle(input.p, input.q)
		delta := math.Abs(input.outRa-iAngle) / input.outRa

		if delta > ppm {
			t.Errorf("fail: want %v %v -> %v got %.6f, %.1f ppm %s",
				input.p,
				input.q,
				input.outRa,
				iAngle,
				delta*1000000,
				input.town,
			)
		}
	}
}

// TestPoleToPole makes sure the sum of internal angles of the north pole, a
// point and the south pole adds up to 180 degrees (pi radians)
func TestPoleToPole(t *testing.T) {
	reindeer := haversine.Coord{90, 0}
	penguins := haversine.Coord{-90, 0}
	for _, input := range tests {
		nAngle := haversine.IntAngle(input.q, reindeer)
		sAngle := haversine.IntAngle(input.q, penguins)
		tAngle := nAngle + sAngle
		delta := math.Abs(tAngle-math.Pi) / math.Pi

		if delta > ppm {
			t.Errorf("fail: want %v -> %-.5f + %-.5f = Pi got %-.5f, %.1f ppm %s",
				input.q,
				nAngle,
				sAngle,
				tAngle,
				delta*1000000,
				input.town,
			)
		}
	}
}
