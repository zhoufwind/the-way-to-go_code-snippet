// Exercise 11.5
// a) Expanding the same program, define a type Triangle, and
// Let it implement AreaInterface. Test this by calculation the area of a specific triangle ( area of a
// Triangle = 0.5 * (base * height) )
// B) Define a new interface PeriInterface which defines a
// Method Perimeter(). Let type Square also implement this interface and test it with our square
// Instance.

// a). 继续扩展程序，定义类型 Triangle，让它实现 AreaInterface 接口。通过计算一个特定三角形的面积来进行测试（三角形面积=0.5 * (底 * 高)）
// b). 定义一个新接口 PeriInterface，它有一个 Perimeter 方法。让 Square 实现这个接口，并通过一个 Square 示例来测试它。

package main

import "fmt"

type Square struct {
    side	float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
    return 4 * sq.side
}

type Triangle struct {
    base	float32
    height	float32
}

func (tr *Triangle) Area() float32 {
    return 0.5 * tr.base * tr.height
}

type AreaInterface interface {
    Area()	float32
}

type PeriInterface interface {
    Perimeter()	float32
}

func main() {
    var areaIntf AreaInterface
    var periIntf PeriInterface

    sq1 := new(Square)
    sq1.side = 5

    tr1 := new(Triangle)
    tr1.base = 3
    tr1.height = 5

    areaIntf = sq1
    fmt.Printf("The square has area: %f\n", areaIntf.Area())

    periIntf = sq1
    fmt.Printf("The square has perimeter: %f\n", periIntf.Perimeter())

    areaIntf = tr1
    fmt.Printf("The triangle has area: %f\n", areaIntf.Area())
}

/*
The square has area: 25.000000
The square has perimeter: 20.000000
The triangle has area: 7.500000
*/
