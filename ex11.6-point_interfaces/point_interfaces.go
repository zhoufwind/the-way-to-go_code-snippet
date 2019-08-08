// Exercise 11.6
// Continue working on the exercise point_methods.go of
// §10.3. Define an interface Magnitude with a function Abs( ). With the methods from §10.3,
// Point, Point3 and Polar implement that interface. Do the same things as in point.go, but now use
// the methods through a variable of the interface type.

// 继续 10.3 中的练习 point_methods.go，定义接口 Magnitude，它有一个方法 Abs()。让 Point、Point3 
// 及Polar 实现此接口。通过接口类型变量使用方法做point.go中同样的事情。

package main

import (
    "fmt"
    "math"
)

type Point struct {
    x	float64
    y	float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

type Point3 struct {
    x, y, z	float64
}

func (p *Point3) Abs() float64 {
    return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2) + math.Pow(p.z, 2))
}

type Polar struct {
    r, t	float64
}

func (p *Polar) Abs() float64 {
    return p.r
}

type Magnitude interface {
    Abs()	float64
}

func main() {
    var magIntf Magnitude

    var p1 = new(Point)
    p1.x = 3
    p1.y = 4
    magIntf = p1
    fmt.Printf("Abs: %f\n", magIntf.Abs())

    var p2 = new(Point3)
    p2.x = 3
    p2.y = 4
    p2.z = 5
    magIntf = p2
    fmt.Printf("Abs: %f\n", magIntf.Abs())

    var p3 = new(Polar)
    p3.r = 1
    magIntf = p3
    fmt.Printf("Abs: %f\n", magIntf.Abs())
}

/*
Abs: 5.000000
Abs: 7.071068
Abs: 1.000000
*/
