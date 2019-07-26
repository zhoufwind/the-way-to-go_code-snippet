// Exercise 10.4
// Define a struct Rectangle with int properties length and width. Give this type methods Area() and
// Perimeter() and test it out.

// 定义一个 Rectangle 结构体，它的长和宽是 int 类型，并定义方法 Area() 和 Primeter()，然后进行测试。

package main

import "fmt"

type Rectangle struct {
    l, w int
}

func Area(rec *Rectangle) (area int) {
    area = rec.l * rec.w
    return
}

func Primeter(rec *Rectangle) (pri int) {
    pri = (rec.l + rec.w) * 2
    return
}

func main() {
    var rec Rectangle
    rec.l = 3
    rec.w = 4
    fmt.Printf("Rectangle's length and width: %v\n", rec)
    fmt.Printf("Area: %d\n", Area(&rec))
    fmt.Printf("Primeter: %d\n", Primeter(&rec))
}

/*
Rectangle's length and width: {3 4}
Area: 12
Primeter: 14
*/
