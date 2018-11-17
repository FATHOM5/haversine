[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](http://godoc.org/github.com/FATHOM5/haversine)
[![Build Status](https://travis-ci.org/FATHOM5/haversine.svg?branch=master)](https://travis-ci.org/FATHOM5/haversine)
[![Go Report Card](https://goreportcard.com/badge/github.com/FATHOM5/haversine?style=flat-square)](https://goreportcard.com/report/github.com/FATHOM5/haversine)
## Haversine

Package haversine provides the great circle distance between two points on the surface of the earth.  Points are identified by longitude and latitude, and distance results can be returned in nautical miles, statute miles, or kilometers.  
![Earth great circle](https://i.imgur.com/iD3X3Ax.png)

The curvature of the earth means that shortest distance between two points cannot be a straight line between the points.  The Haversine formula provides an approximation for the shortest great circle route over the surface of earth that connects the two points.  In this figure the dotted yellow line is the arc of a great circle. *Image courtesy USGS.*

## Installation

`go get github.com/FATHOM5/haversine`

## Usage

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
    }

## Documenation

* http://godoc.org/github.com/FATHOM5/haversine

## References

* Many thanks to the original author, [umahmood](https://github.com/umahmood/haversine), for a great package to begin this fork.
* https://plus.maths.org/content/lost-lovely-haversine
* https://en.wikipedia.org/wiki/Haversine_formula

## License

See [LICENSE](LICENSE.md) for rights and limitations (MIT).
