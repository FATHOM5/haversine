// Example demonstrates use of the haversine Distance function.
package haversine_test

import (
	"fmt"

	"github.com/loraxipam/haversine"
)

func ExampleDistanceNM() {
	austin := haversine.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := haversine.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}

	nm := haversine.DistanceNM(austin, paloAlto)

	fmt.Printf("%.1f miles is a long walk to Silicon Valley.\n", nm)

	// Output: 1286.1 miles is a long walk to Silicon Valley.
}

func ExampleDistance() {
	austin := haversine.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := haversine.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}

	km := haversine.Distance(austin, paloAlto, 58230)

	fmt.Printf("%.1f km is a long walk to Silicon Valley on Saturn.\n", km)

	// Output: 21770.2 km is a long walk to Silicon Valley on Saturn.
}

func ExampleIntAngle() {
	austin := haversine.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := haversine.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: -122.1430,
	}

	radians := haversine.IntAngle(austin, paloAlto)

	fmt.Printf("%.4f radians is a long walk to Silicon Valley, perhaps.\n", radians)

	// Output: 0.3739 radians is a long walk to Silicon Valley, perhaps.
}

func ExampleNmToKm() {
	fmt.Printf("One degree is 60 NM which is %.1f kilometers\n", haversine.NmToKm(60))

	// Output: One degree is 60 NM which is 111.1 kilometers
}

func ExampleDegrees() {
	// a planety thing
	type planet struct {
		name string
		r    float64
	}

	// km radius
	var solarSystem = []planet{
		{"the Sun", 695700},
		{"the Moon", 1737},
		{"Mercury", 2440},
		{"Venus", 6051},
		{"Earth", 6371},
		{"Mars", 3390},
		{"Jupiter", 69910},
		{"Saturn", 58230},
		{"Uranus", 25360},
		{"Neptune", 24620},
		{"the unit sphere", 1},
		// Sorry, Pluto, you got kicked to the curb, dude.
		// Your sorry ass is even smaller than the moon, anyway, so,
		// "Meh" is what I say to you. Sorry, Clyde. I love you, man.
	}

	// a winter seaside resort town
	wilbur := haversine.Coord{29.13, -80.97}
	fmt.Println("Wilbur-by-the-Sea is here:", wilbur)

	b := wilbur.Radians()
	fmt.Println(" which, in radians, is", b)
	fmt.Println(" and inverted back to degrees yields", b.Degrees())

	// it's antipode, a not winter not seaside not resort town
	summerhills := haversine.Coord{45.71, -122.43}
	fmt.Println("Summer Hills, however, is here:", summerhills)

	// regardless of sphere, what's the angle?
	fmt.Printf(" That's an angular distance of %.4f radians\n", haversine.IntAngle(wilbur, summerhills))
	fmt.Printf(" or %.2f degrees\n", 180/3.14159*haversine.IntAngle(wilbur, summerhills))
	fmt.Printf(" so from Wilbur to Summer Hills is %.0f miles.\n", haversine.DistanceMi(wilbur, summerhills))

	// if this were a solar system object
	for _, v := range solarSystem {
		fmt.Printf("On %s, that would be %.1f miles or %.1f days of driving\n", v.name, haversine.Distance(wilbur, summerhills, haversine.NmToMi(haversine.KmToNm(v.r))), 6*v.r/6371)
	}

	// Output:
	// Wilbur-by-the-Sea is here: {29.13 -80.97}
	//  which, in radians, is {0.5084144111059482 -1.4131930953398086}
	//  and inverted back to degrees yields {29.130000000000003 -80.97}
	// Summer Hills, however, is here: {45.71 -122.43}
	//  That's an angular distance of 0.6342 radians
	//  or 36.33 degrees
	//  so from Wilbur to Summer Hills is 2510 miles.
	// On the Sun, that would be 274083.4 miles or 655.2 days of driving
	// On the Moon, that would be 684.3 miles or 1.6 days of driving
	// On Mercury, that would be 961.3 miles or 2.3 days of driving
	// On Venus, that would be 2383.9 miles or 5.7 days of driving
	// On Earth, that would be 2510.0 miles or 6.0 days of driving
	// On Mars, that would be 1335.6 miles or 3.2 days of driving
	// On Jupiter, that would be 27542.3 miles or 65.8 days of driving
	// On Saturn, that would be 22940.7 miles or 54.8 days of driving
	// On Uranus, that would be 9991.0 miles or 23.9 days of driving
	// On Neptune, that would be 9699.5 miles or 23.2 days of driving
	// On the unit sphere, that would be 0.4 miles or 0.0 days of driving

}
