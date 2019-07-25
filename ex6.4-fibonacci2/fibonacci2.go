// Exercise 6.4
// Rewrite the Fibonacci program above to return 2 named variables (see § 6.2),
// namely the value and the position of the Fibonacci-number, like 5 and 4 or 89
// and 10.

// 重写本节中生成斐波那契数列的程序并返回两个命名返回值（详见第 6.2 节），即数列中的位置和对应的值，例如 5 与 4，89 与 10。

package main

import "fmt"

func main() {
    res, pos := fibonacci(4)
    fmt.Printf("Position: %d, value: %d\n", pos, res)
    res, pos = fibonacci(10)
    fmt.Printf("Position: %d, value: %d\n", pos, res)
}

func fibonacci(n int) (res int, pos int) {
    if n <= 1 {
        res = 1
    } else {
        res1, _ := fibonacci(n-1)
        res2, _ := fibonacci(n-2)
        res = res1 + res2
    }
    pos = n
    return
}

/*
Position: 4, value: 5
Position: 10, value: 89
*/
