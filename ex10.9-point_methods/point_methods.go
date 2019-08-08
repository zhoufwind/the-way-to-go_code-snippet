// Exercise 10.9
// Start from point.go ( exercise from § 10.1): now implement the functions Abs()
// and Scale() as methods with the receiver type Point. Also implement Abs() for
// Point3 and Polar as methods. Do the same things as in point.go, but now use the
// methods.

// 从 point.go 开始（第 10.1 节的联系）：使用方法来实现 Abs() 和 Scale()函数，
// Point 作为方法的接收者类型。也为 Point3 和 Polar 实现 Abs() 方法。完成了 
// point.go 中同样的事情，只是这次通过方法。

package main

import (
    "fmt"
    "math"
)

type Point struct {
    x, y	float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) Scale(factor float64) {
    p.x *= factor
    p.y *= factor
}

type Point3 struct {
    x, y, z float64
}

func (p *Point3) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

type Polar struct {
    r, t	float64
}

func (p Polar) Abs() float64 { return p.r }

func main() {
    p := Point{3.0, 4.0}
    fmt.Println(p.Abs())

    p.Scale(5.0)
    fmt.Printf("%v\n", p)

    p2 := new(Point3)
    p2.x = 3.0
    p2.y = 4.0
    p2.z = 5.0
    fmt.Println(p2.Abs())
}

/*
5
{15 20}
7.0710678118654755
*/
