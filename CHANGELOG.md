# Change Log
All notable changes documented in adherance to [Semantic Versioning](http://semver.org/).

v1.0.0 - Original API last committed in Nov 2015

v2.0.0 - API changes committed in Nov 2018.  See commit [092bd4b](https://github.com/FATHOM5/haversine/commit/092bd4b5508f4355bafeb2d632eded2cf958a251)

v3.0.0 - API changes committed in May 2020.  See commit [](https://github.com/FATHOM5/haversine/commit/somenumbergoeshere)
  Refactored Distance to use a third input param, the radius of the sphere, to create a more generalized function.
  Created DistanceNM(p, q Coord) function to return nautical miles on Earth.
  Created IntAngle(p, q Coord) function to return radians between two points. Above Distance function calls this.
  Exported miles, kilometer and nautical mile Earth radius values
  Exported distance conversion functions
  Fixed errors in haversine_test.go
