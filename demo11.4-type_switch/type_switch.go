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

    switch t := areaI.(type) {
    case *Square:
        fmt.Printf("Type Square %T with value %v\n", t, t)
    case *Circle:
        fmt.Printf("Type Circle %T with value %v\n", t, t)
    case nil:
        fmt.Printf("nil value: nothing to check?\n")
    default:
        fmt.Printf("Unexpected type %T\n", t)
    }
}

/*
Type Square *main.Square with value &{5}
*/
