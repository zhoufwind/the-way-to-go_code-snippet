// Exercise 7.3
// In § 6.6 we saw a recursive solution for calculating Fibonacci
// numbers. But they can also be calculated in an imperative way, using a simple
// array. Do this for the first 50 Fibonacci numbers.

// 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。

package main

import "fmt"
import "time"

func main() {
    start := time.Now()

    var f [46]int
    f[0] = 1
    f[1] = 1
    for i := 2; i < len(f); i++ {
        f[i] = f[i-1] + f[i-2]
    }
    for ix, v := range f {
        fmt.Printf("fibonacci(%d) is: %d\n", ix, v)
    }

    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("longCalculation took this amount of time: %s\n", delta)
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
longCalculation took this amount of time: 145.638µs
*/
