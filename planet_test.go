package haversine_test

import (
        "fmt"

        "github.com/FATHOM5/haversine"
)

func main(){
        // a planety thing
        type planet struct {
                name string
                r float64
        }

        // km radius
        var solarSystem = []planet {
                {"the Sun",695700},
                {"the Moon",1737},
                {"Mercury",2440},
                {"Venus",6051},
                {"Earth",6371},
                {"Mars",3390},
                {"Jupiter",69910},
                {"Saturn",58230},
                {"Uranus",25360},
                {"Neptune",24620},
                {"the unit sphere",1},
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
        fmt.Printf("That's an angular distance of %.4f radians\n", haversine.IntAngle(wilbur, summerhills))
        fmt.Printf(" or %.2f degrees\n", 180 / 3.14159 * haversine.IntAngle(wilbur, summerhills))
        fmt.Printf(" so from Wilbur to Summer Hills is %.0f miles.\n", haversine.DistanceMi(wilbur, summerhills))

        // if this were a solar system object
        for _, v := range solarSystem {
                fmt.Printf("On %s, that would be %.1f miles or %.1f days of driving\n", v.name, haversine.Distance(wilbur, summerhills, v.r / 1.6093), 6*v.r/6371)
        }

}
