// Exercise 7.9
// Given a slice s []int and a factor of type int, enlarge s so that its
// new length is len(s) * factor.

// 给定 slice s[]int 和一个 int 类型的因子，扩展 s 使其长度为 len(s) * factor。

package main

import "fmt"

func main() {
    s := make([]int, 2, 10)
    fmt.Printf("array(%v) length: %d, capacity: %d\n", s, len(s), cap(s))

    factor := 6
    fmt.Printf("factor: %d\n", factor)

    m := len(s)
    n := m * factor
    if n > cap(s) {
        fmt.Printf("slice bounds out of range, reslice!\n")
        ns := make([]int, (n+1)*2)
        copy(ns, s)
        s = ns
    }
    s = s[0:n]
    fmt.Printf("array(%v) length: %d, capacity: %d\n", s, len(s), cap(s))
}

/*
array([0 0]) length: 2, capacity: 10
factor: 6
slice bounds out of range, reslice!
array([0 0 0 0 0 0 0 0 0 0 0 0]) length: 12, capacity: 26
*/
