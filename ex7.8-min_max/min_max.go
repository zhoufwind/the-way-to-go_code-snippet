// Exercise 7.8
// Write a minSlice function which takes a slice of ints and returns the minimum,
// and a maxSlice function which takes a slice of ints and returns the maximum.

// 写一个 minSlice 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice 方法返回最大值。

package main

import "fmt"

func main() {
    arr1 := []int{5, 0, 3, -1, 2}
    sl1 := arr1[:]
    fmt.Printf("array(%v) min: %v, max: %v\n", sl1, minSlice(sl1), maxSlice(sl1))
}

func minSlice(a []int) (min int) {
    min = a[0]
    for ix, _ := range a {
        if min > a[ix] {
            min = a[ix]
        }
    }
    return
}

func maxSlice(a []int) (max int) {
    max = a[0]
    for ix, _ := range a {
        if max < a[ix] {
            max = a[ix]
        }
    }
    return
}

/*
array([5 0 3 -1 2]) min: -1, max: 5
*/
