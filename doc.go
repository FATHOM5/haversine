/*
Package haversine provides the great circle distance between two points
on the surface of the earth.  Points are identified by longitude and
latitude, and distance results are returned in nautical miles.  Functions
are also provided to return the distance in kilometers or statute miles.

The below example shows how to calculate the shortest path between two
coordinates on the surface of the Earth.

    package main

    import (
        "fmt"

        "github.com/FATHOM5/haversine"
    )

    func main() {
        oxford := haversine.Coord{Lat: 51.45, Lon: 1.15}  // Oxford, UK
        turin  := haversine.Coord{Lat: 45.04, Lon: 7.42}  // Turin, Italy

        nm := haversine.Distance(oxford, turin)

        fmt.Println("Nautical Miles:", nm)

        // Similarly, distance can be obtained in other common units
        km := haversine.DistanceKm(oxford, turin)
        fmt.Println("Kilometers:", km)
        mi := haversine.DistanceMi(oxford, turin)
        fmt.Println("Statute Miles:", mi)
    }
*/
package haversine
