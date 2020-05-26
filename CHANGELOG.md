# Change Log
All notable changes documented in adherance to [Semantic Versioning](http://semver.org/).

Prior to June 2020 the API for this package and it's parent at umahmood/haversine had major version number changes for each  adaptation of the API.  In June 2020, the original 2015 API was renumbered to v0.1.0 to reflect that it was premature to declare a v1.0.0 release without a larger user community for this small package. Edits to the package in 2018 by Fathom5 and open source contributions in 2020 were considered minor API version changes prior to an official 1.0.0 release that will occur when the package is used by at least ten other projects or contributors.

This change will better reflect the package's maturity in the [golang package ecosystem](https://pkg.go.dev/)

v0.1.0 - Original API last committed in Nov 2015

v0.2.0 - API changes committed in Nov 2018.  See commit [092bd4b](https://github.com/FATHOM5/haversine/commit/092bd4b5508f4355bafeb2d632eded2cf958a251)

v0.3.0 - API changes committed in May 2020. See commit [](https://github.com/FATHOM5/haversine/commit/.......)
  Refactored Distance to use a third input param, the radius of the sphere, to create a more generalized function.
  Created DistanceNM(p, q Coord) function to return nautical miles on Earth.
  Created IntAngle(p, q Coord) function to return radians between two points. Above Distance function calls this.
  Exported miles, kilometer and nautical mile Earth radius values
  Exported distance conversion functions
  Fixed errors in haversine_test.go
