// Exercise 6.6
// Write a program which prints the factorial (!) of the first 30 integers
// The factorial n! of a number n is defined as: n! = n * (n-1)!, 0!=1
// So this clearly a good candidate for a recursive function Factorial.
// Make a 2nd version of Factorial with a named return variable.
// Remark that when using type int the calculation is only correct up until 12!, this
// is of course because an int can only contain integers which fit in 32 bit. Go doesn’t
// warn against this overflow-error! How can you gset more correct results ?
// The best solution is to use the big package (see § 9.4).

// 实现一个输出前 30 个整数的阶乘的程序。
// n! 的阶乘定义为：n! = n * (n-1)!, 0! = 1，因此它非常适合使用递归函数来实现。
// 然后，使用命名返回值来实现这个程序的第二个版本。
// 特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误。那么，如何才能解决这个问题呢？
// 最好的解决方案就是使用 big 包（详见第 9.4 节）。

package main

import "fmt"

func main() {
    for i := uint64(0); i < uint64(30); i++ {
        fmt.Printf("%d! = %d\n", i, Factorial(i))
    }
}

func Factorial(n uint64) (res uint64) {
    if n > 0 {
        res = n * Factorial(n-1)
        return res
    } else {
        return 1
    }
}

/*
0! = 1
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
6! = 720
7! = 5040
8! = 40320
9! = 362880
10! = 3628800
11! = 39916800
12! = 479001600
13! = 6227020800
14! = 87178291200
15! = 1307674368000
16! = 20922789888000
17! = 355687428096000
18! = 6402373705728000
19! = 121645100408832000
20! = 2432902008176640000
21! = 14197454024290336768
22! = 17196083355034583040
23! = 8128291617894825984
24! = 10611558092380307456
25! = 7034535277573963776
26! = 16877220553537093632
27! = 12963097176472289280
28! = 12478583540742619136
29! = 11390785281054474240
*/
