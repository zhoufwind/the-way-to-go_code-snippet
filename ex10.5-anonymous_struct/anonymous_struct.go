// Exercise 10.5
// Make a struct with 1 named float field, and 2 anonymous fields of type int and
// string. Create an instance of the struct with a literal expression and print the
// content.

// 创建一个结构体，它有一个具名的 float 字段，2 个匿名字段，类型分别是 int 和 string。
// 通过结构体字面量新建一个结构体实例并打印它的内容。

package main

import "fmt"

type struct1 struct {
    f1	float32
    int
    string
}

func main() {
    s1 := new(struct1)
    s1.f1 = 5.5
    s1.int = 10
    s1.string = "Stephen"

    fmt.Printf("Float: %f\n", s1.f1)
    fmt.Printf("Int: %d\n", s1.int)
    fmt.Printf("String: %s\n", s1.string)

    s2 := struct1{3.3, 20, "Zac"}
    fmt.Println(s2)
}

/*
Float: 5.500000
Int: 10
String: Stephen
{3.3 20 Zac}
*/
