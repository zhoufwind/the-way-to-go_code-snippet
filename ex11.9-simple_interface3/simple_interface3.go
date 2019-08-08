// Exercise 11.9
// Continuing on Exercise 11.2, add a function gI which, instead of an Simpler type, accepts a
// empty interface parameter. The function then tests with a type assertion if the parameter fulfills
// the Simpler type. Now call this gI function instead of the fI function. Make your code as safe as
// possible.

// 继续练习11.2，在它中添加一个 gI 函数，它不再接受 Simpler 类型的参数，而是接受一个空接口参数。
// 然后通过类型断言判断参数是否是 Simpler 类型。最后在 main 使用 gI 取代 fI 函数并调用它。确保你的代码足够安全。

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
    i, j	int
}

func (rs *RSimple) Get() int {
    return rs.j
}

func (rs *RSimple) Set(n int) {
    rs.j = n
}

type Simpler interface {
    Get() int
    Set(int)
}

func fI(it Simpler) int {
    switch it.(type) {
    case *Simple:
        it.Set(5)
        return it.Get()
    case *RSimple:
        it.Set(50)
        return it.Get()
    default:
        return 99
    }
    return 0
}

func gI(any interface{}) int {
    // return any.(Simpler).Get()	// unsafe, runtime panic possible
    if v, ok := any.(Simpler); ok {
        return v.Get()
    }
    return 0	// default value
}

func main() {
    var s Simple = Simple{6}
    fmt.Println(gI(&s))	// &s is required because Get() is defined with a receiver type pointer

    var rs RSimple = RSimple{60, 60}
    fmt.Println(gI(&rs))
}

/*
6
60
*/
