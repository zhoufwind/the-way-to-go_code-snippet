// Exercise 10.3
// Define a 2 dimensional Point with coordinates X and Y as a struct. Do the same
// for a 3 dimensional point, and a Polar point defined with its polar coordinates.
// Implement a function Abs() that calculate the length of the vector represented by
// a Point, and a function Scale that multiplies the coordinates of a point with a scale
// factor(hint: use function Sqrt from package math).

// 使用坐标 X、Y 定义一个二维 Point 结构体。同样地，对一个三维点使用它的极坐标定义一个 
// Polar 结构体。实现一个Abs() 方法来计算一个 Point 表示的向量的长度，实现一个 Scale 
// 方法，它将点的坐标乘以一个尺度因子（提示：使用math 包里的 Sqrt 函数）
//（function Scale that multiplies the coordinates of a point with a scale factor）。

package main

import (
    "fmt"
    "math"
)

type Point struct {
    x	float64
    y	float64
}

func Abs(p *Point, q *Point)(res float64) {
    res = math.Sqrt(math.Pow((p.x-q.x),2)+math.Pow((p.y-q.y),2))
    return res
}

func Scale(p *Point, factor float64) {
    p.x = p.x * factor
    p.y = p.y * factor
}

func main() {
    var a,b Point

    a.x = 6
    a.y = 5

    b.x = 2
    b.y = 2

    fmt.Printf("a value: %v, b value: %v\n", a, b)
    fmt.Printf("Distance Vector: %v\n", Abs(&a, &b))

    Scale(&a, 5)
    fmt.Println(&a)
}

/*
a value: {6 5}, b value: {2 2}
Distance Vector: 5
&{30 25}
*/
