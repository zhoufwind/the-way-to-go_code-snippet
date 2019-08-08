// Exercise 11.1
// Define an interface Simpler with methods Get() which returns an integer, and Set() which has an
// integer as parameter. Make a struct type Simple which implements this interface.
// Then define a function which takes a parameter of the type Simpler and calls both methods upon
// it. Call this function from main to see if it all works correctly.

// 定义一个接口 Simpler，它有一个 Get() 方法和一个 Set()，Get()返回一个整型值，Set() 有一个整型参数。创建一个结构体类型 Simple 实现这个接口。
// 接着定一个函数，它有一个 Simpler 类型的参数，调用参数的 Get() 和 Set() 方法。在 main 函数里调用这个函数，看看它是否可以正确运行。

package main

import "fmt"

type Simple struct {
    int
}

func (s *Simple) Get() int {
    return s.int
}

func (s *Simple) Set(n int) {
    s.int = n
}

type Simpler interface {
    Get()	int
    Set(int)
}

func IntfOps(si Simpler) int {
    si.Set(15)
    return si.Get()
}

func main() {
    s := &Simple{5}
    fmt.Println(s.Get())
    s.Set(10)
    fmt.Println(s.Get())

    fmt.Println(IntfOps(s))
}

/*
5
10
15
*/
