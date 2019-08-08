package main

import (
    "fmt"
    "math"
)

type Square struct {
    side	float32
}

func (s *Square) Area() float32 {
    return s.side * s.side
}

type Circle struct {
    radius	float32
}

func (c *Circle) Area() float32 {
    return math.Pi * c.radius * c.radius
}

type Shaper interface {
    Area()	float32
}

func main() {
    var areaI Shaper

    sq1 := new(Square)
    sq1.side = 5

    areaI = sq1
    // Is Square the type of areaI?
    if t, ok := areaI.(*Square); ok {
        fmt.Printf("The type of areaI is: %T\n", t)
    }
    if u, ok := areaI.(*Circle); ok {
        fmt.Printf("The type of areaI is: %T\n", u)
    } else {
        fmt.Println("areaI does not contain a variable of type Circle")
    }
}

/*
The type of areaI is: *main.Square
areaI does not contain a variable of type Circle
*/
