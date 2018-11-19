// Example demonstrates use of the haversine Distance function.
package haversine_test

import (
	"fmt"

	"github.com/FATHOM5/haversine"
)

func Example() {
	austin := haversine.Coord{ // Austin, Texas
		Lat: 30.2672,
		Lon: -97.7431,
	}
	paloAlto := haversine.Coord{ // Palo Alto, California
		Lat: 37.4419,
		Lon: 122.1430,
	}

	nm := haversine.Distance(austin, paloAlto)

	fmt.Printf("%.1f miles is a long walk to Silicon Valley.\n", nm)
}
