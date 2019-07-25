package main

import "fmt"
import "math"

func mySqrt(f float64) (v float64, ok bool) {
    if f < 0 { return }
    return math.Sqrt(f), true
}

func main() {
    t, ok := mySqrt(9.0)
    if ok {
        fmt.Println(t)
    }
}
