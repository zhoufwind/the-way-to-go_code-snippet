// Exercise 9.4
// Make a program main_oddeven.go which tests for the first 100 integers whether
// they are even or not. The function which does the test is contained in a package
// even.

// 创建一个程序 main_oddven.go 判断前 100 个整数是不是偶数，包内同时包含测试的功能。

package main

import (
    "fmt"
    "./even"
)

func main() {
    for i := 0; i < 100; i++ {
        fmt.Printf("Is the integer %d even? %t\n", i, even.Even(i))
    }
}

/*
Is the integer 0 even? true
Is the integer 1 even? false
Is the integer 2 even? true
Is the integer 3 even? false
Is the integer 4 even? true
...
Is the integer 96 even? true
Is the integer 97 even? false
Is the integer 98 even? true
Is the integer 99 even? false
*/
