// Exercise 9.2
// Use a function from the unsafe package to test the size of an int variable on your computer.

// 通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。

// 【golang】unsafe.Sizeof浅析
// Ref: https://blog.csdn.net/haodawang/article/details/80005072

package main

import "fmt"
import "unsafe"

func main() {
    var i int = 100
    fmt.Printf("The size of an int is: %d\n", unsafe.Sizeof(i))

    slice := []int{1,2,3,4}
    fmt.Printf("The size of a slice is: %d\n", unsafe.Sizeof(slice))

    arr := [...]int{1,2,3,4}
    fmt.Printf("The size of an array(arr) is: %d\n", unsafe.Sizeof(arr))

    arr2 := [...]int{1,2,3,4,5}
    fmt.Printf("The size of an array(arr2) is: %d\n", unsafe.Sizeof(arr2))
}

/*
The size of an int is:  8
*/
