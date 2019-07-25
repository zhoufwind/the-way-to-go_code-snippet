// Exercise 7.14
// Write a program that reverses a string, so “Google” is printed as ” elgooG”.
// (Hint: use a slice of bytes and conversions.)
// If you coded a solution with two slices, try a variant which uses only one (Hint:
// use swapping)
// If you want to be able to reverse Unicode-strings: use [ ]int !

// 编写一个程序，要求能够反转字符串，即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。
// 如果您使用两个切片来实现反转，请再尝试使用一个切片（提示：使用交换法）。
// 如果您想要反转 Unicode 编码的字符串，请使用 []int 类型的切片。

package main

import "fmt"

func main() {
    a := "google"
    fmt.Printf("%s reverse: %s\n", a, StringReverse(a))
}

func StringReverse(a string) (b string) {
    sl := []byte(a)
    sl2 := make([]byte, len(sl))
    for ix, value := range sl {
        sl2[len(sl)-1-ix] = value
    }
    b = string(sl2)
    return
}

/*
google reverse: elgoog

Comment: not suppport Chinese
e.g.: 你好吗 reverse: ??彥堽?
*/
