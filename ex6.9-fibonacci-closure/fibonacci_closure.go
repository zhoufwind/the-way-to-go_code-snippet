// Exercise 6.9
// Write a non-recursive version of the Fibonacci program from § 6.6 using a function
// as a closure

// 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序。

package main

import "fmt"
import "time"

func main() {
    start := time.Now()
    f := fibonacci()
    for i := 0; i < 46; i++ {
        fmt.Printf("fibonacci(%d) is: %d\n", i, f())
    }
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
    previous := 0
    current := 0
    t := 0
    return func() int {
        if current == 0 {
            previous = 0
            current = 1
        } else {
            t = previous
            previous = current
            current = t + current
        }
        return current
    }
}

/*
fibonacci(0) is: 1
fibonacci(1) is: 1
fibonacci(2) is: 2
fibonacci(3) is: 3
fibonacci(4) is: 5
fibonacci(5) is: 8
fibonacci(6) is: 13
fibonacci(7) is: 21
fibonacci(8) is: 34
fibonacci(9) is: 55
fibonacci(10) is: 89
fibonacci(11) is: 144
fibonacci(12) is: 233
fibonacci(13) is: 377
fibonacci(14) is: 610
fibonacci(15) is: 987
fibonacci(16) is: 1597
fibonacci(17) is: 2584
fibonacci(18) is: 4181
fibonacci(19) is: 6765
fibonacci(20) is: 10946
fibonacci(21) is: 17711
fibonacci(22) is: 28657
fibonacci(23) is: 46368
fibonacci(24) is: 75025
fibonacci(25) is: 121393
fibonacci(26) is: 196418
fibonacci(27) is: 317811
fibonacci(28) is: 514229
fibonacci(29) is: 832040
fibonacci(30) is: 1346269
fibonacci(31) is: 2178309
fibonacci(32) is: 3524578
fibonacci(33) is: 5702887
fibonacci(34) is: 9227465
fibonacci(35) is: 14930352
fibonacci(36) is: 24157817
fibonacci(37) is: 39088169
fibonacci(38) is: 63245986
fibonacci(39) is: 102334155
fibonacci(40) is: 165580141
fibonacci(41) is: 267914296
fibonacci(42) is: 433494437
fibonacci(43) is: 701408733
fibonacci(44) is: 1134903170
fibonacci(45) is: 1836311903
longCalculation took this amount of time: 133.553µs
*/
