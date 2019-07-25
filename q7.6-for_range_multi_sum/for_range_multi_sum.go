// Question 7.6
// Sum up the contexts in the Go syntax where the ellipsis operator … is used.

// 通过使用省略号操作符 ... 来实现累加方法。

package main

import "fmt"

func main() {
    fmt.Println(Sum(1, 2, 3, 4, 5))

    arr := []int{6, 7, 8, 9, 10}
    fmt.Println(Sum(arr...))
}

func Sum(a ...int) (sum int) {
    for _, v := range a {
        sum += v
    }
    return
}

/*
15
40
*/
