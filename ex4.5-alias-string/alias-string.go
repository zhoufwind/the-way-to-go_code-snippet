// Exercise 4.5
// Define an alias type Rope for string and declare a variable with it.

// 定义一个 string 的类型别名 Rope，并声明一个该类型的变量。

package main

import "fmt"

type Rope string

func main() {
    var a Rope = "hello world"
    fmt.Println(a)
}
