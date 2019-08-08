// Exercise 10.12
// Given a struct type T: type T struct {
//  a int
//  b float32
//  c string
// }
//  and a value t: t := &T{ 7, -2.35, “abc\tdef” }
//  make a String() method for T so that for fmt.Printf(“%v\n”, t)
//  the following format is send to the output: 7 / -2.350000 / “abc\tdef”

// 给定结构体类型 T:
// type T struct {
//     a int
//     b float32
//     c string
// }
// 值 t: t := &T{7, -2.35, "abc\tdef"}。给 T 定义 String()，使得 fmt.Printf("%v\n", t) 输出：7 / -2.350000 / "abc\tdef"。

package main

import (
    "fmt"
    //"strconv"
)

type T struct {
    a	int
    b	float64
    c	string
}

func (t *T) String() string {
    //return strconv.Itoa(t.a) + " / " + strconv.FormatFloat(t.b, 'E', -1, 64) + " / \"" + t.c + "\""
    return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}

func main() {
    t := &T{7, -2.35, "abc\tdef"}
    fmt.Println(t)
}

/*
7 / -2.350000 / "abc\tdef"
*/
