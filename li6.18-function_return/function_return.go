package main

import "fmt"

func main() {
    var num int = 3

    // make an Add2 function, give it a name p2, and call it:
    // 函数 Add2 和 Adder 均会返回签名为 func(b int) int 的函数
    p2 := Add2()
    fmt.Printf("Call Add2 for %d gives: %v\n", num, p2(num))

    // make a special Adder function, a gets value 3:
    TwoAdder := Adder(2)
    fmt.Printf("Call Adder(2) for %d gives: %v\n", num, TwoAdder(num))
}

func Add2() func(b int) int {
    return func(b int) int {
        return b + 2
    }
}

func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}

/*
Call Add2 for 3 gives: 5
Call Adder(2) for 3 gives: 5
*/
