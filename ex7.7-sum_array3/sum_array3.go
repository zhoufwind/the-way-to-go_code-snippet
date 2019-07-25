// Exercise 7.7
// Write a function SumAndAverage which returns these two as unnamed variables
// of type int and float32.

// 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。

package main

import "fmt"

func main() {
    a := []int{1, 2, 3}
    b := []float64{1.1, 2.2, 3.3}
    sla, slb := SumAndAverage(a, b)
    fmt.Printf("Sum: %f, avg: %g\n", sla, slb)
}

func SumAndAverage(a []int, b []float64) (sum float64, avg float64) {
    for _, va := range a {
        sum += float64(va)
    }
    for _, vb := range b {
        sum += vb
    }
    avg = sum / float64((len(a) + len(b)))
    return
}

/*
Sum: 12.600000, avg: 2.1
*/
