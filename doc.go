/*
Package haversine provides the great circle distance between two points
on a sphere. Several functions return great circle distances of the Earth.
Points are identified by latitude and longitude in degrees, and distance
results are returned in nautical miles.  Functions are also provided to
return the distance in kilometers or statute miles, or the geocentric
angular separation between those points.

The below example shows how to calculate the shortest path between two
coordinates on the surface of the Earth.

    package main

    import (
        "fmt"
        "math"

        "github.com/loraxipam/haversine"
    )

    func main() {
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
        fmt.Printf("it's only %.1f miles on the moon, though.\n", haversine.Distance(austin, paloAlto, 937.9))
        fmt.Printf("That's an angle of %.1f degrees\n", 180/math.Pi*haversine.IntAngle(austin, paloAlto))
    }
*/
package haversine
