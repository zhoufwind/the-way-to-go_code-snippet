// Exercise 11.4
// Continuing with exercise 11.1, make a second type RSimple which also implements the interface
// Simpler.
// Expand the function fI so that it can distinguish between variables of both types.

// 接着练习 11.1 中的内容，创建第二个类型 RSimple，它也实现了接口 Simpler，写一个函数 fi，使
// 它可以区分Simple 和 RSimple 类型的变量。

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

type RSimple struct {
    int
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

    var siI Simpler
    siI = s

    if t, ok := siI.(*Simple); ok {
        fmt.Printf("The type of siI is: %T\n", t)
    }
}

/*
5
The type of siI is: *main.Simple
*/
