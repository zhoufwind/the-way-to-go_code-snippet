// Exercise 6.3
// Make a function that has as arguments a variable number of ints and which prints
// each integer on a separate line.

// 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。

package main

import "fmt"

func main() {
    vararg(1, 3, 5, 7, 9, 99)
}

func vararg(a ...int) {
    if len(a) == 0 {
        return
    }
    for _, v := range a {
        fmt.Println(v)
    }
}

/*
1
3
5
7
9
99
*/
