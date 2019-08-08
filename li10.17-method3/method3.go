package main

import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
    Point
    name	string
}

func main() {
    n := &NamedPoint{Point{3, 4}, "Pythagoras"}
    fmt.Println(n.Abs())

    n2 := NamedPoint{Point{6, 8}, "P2"}
    fmt.Println(n2.Abs())
}

/*
5
10
*/
