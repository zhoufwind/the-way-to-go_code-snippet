package main

import "fmt"

func main() {
    arr1 := [4] float64 {1.1, 2.2, 3.3, 4.4}
    sl := arr1[:]
    fmt.Println(Sum(sl))
}

func Sum(a []float64) (sum float64) {
    for _, v := range a {
        sum += v
    }
    return
}
