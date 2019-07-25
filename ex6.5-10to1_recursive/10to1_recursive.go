// Exercise 6.5
// Print the numbers from 10 to 1 in that order using a recursive function printrec(i
// int)

// 使用递归函数从 10 打印到 1。

package main

import "fmt"

func main() {
    PrintNum(1)
}

func PrintNum(i int) {
    if i > 10 {
        return
    }
    PrintNum(i + 1)
    fmt.Printf("%d\n", i)
}

/*
10
9
8
7
6
5
4
3
2
1
*/
