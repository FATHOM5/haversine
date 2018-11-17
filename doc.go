/*
Package haversine provides the great circle distance between two points
on the surface of the earth.  Points are identified by latitude and
longitude, and distance results are returned in nautical miles.  Functions
are also provided to return the distance in kilometers or statute miles.

The below example shows how to calculate the shortest path between two
coordinates on the surface of the Earth.

    package main

    import (
        "fmt"

        "github.com/FATHOM5/haversine"
    )

    func main() {
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
*/
package haversine
