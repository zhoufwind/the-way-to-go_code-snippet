package main

import "fmt"

// 关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值
func f() (ret int) {
    defer func() {
        ret++
    }()
    return 1
}

// This can be convenient for modifying the error return value of a function.
// 可用于在返回语句之后修改返回的 error 时使用
func main() {
    fmt.Println(f())
}

/*
2
*/
