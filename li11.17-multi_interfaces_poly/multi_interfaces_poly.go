package main

import "fmt"

type Shaper interface {
    Area()	float32
}

type TopologicalGenus interface {
    Rank()	int
}

type Square struct {
    side	float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (sq *Square) Rank() int {
    return 1
}

type Rectangle struct {
    length, width	float32
}

func (r Rectangle) Area() float32 {
    return r.length * r.width
}

func (r Rectangle) Rank() int {
    return 2
}

func main() {
    r := Rectangle{5, 3}	// Area() of Rectangle needs a value
    q := &Square{5}	// Area() of Square needs a pointer

    //shapes := []Shaper{Shaper(r), Shaper(q)}

    shapes := []Shaper{r, q}

    fmt.Println("Looping through shapes for area ...")
    for i, _ := range shapes {
        fmt.Println("Shape details: ", shapes[i])
        fmt.Println("Area of this shape is: ", shapes[i].Area())
    }

    topgen := []TopologicalGenus{r, q}

    fmt.Println("Looping through topgen for rank ...")
    for i, _ := range topgen {
        fmt.Println("Shape details: ", topgen[i])
        fmt.Println("Topological Genus of this shape is: ", topgen[i].Rank())
    }
}

/*
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25
Looping through topgen for rank ...
Shape details:  {5 3}
Topological Genus of this shape is:  2
Shape details:  &{5}
Topological Genus of this shape is:  1
*/
