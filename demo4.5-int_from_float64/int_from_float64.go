package main

import (
    "fmt"
    "math"
)

func main() {
    var a float64 = 2147483647.1
    b := IntFromFloat64(a)
    fmt.Println(b)
}

func IntFromFloat64(x float64) int {
    if math.MinInt32 <= x && x <= math.MaxInt32 {	// x lies in the integer range
        whole, fraction := math.Modf(x)
        if fraction >= 0.5 {
            whole++
        }
        return int(whole)
    }
    panic(fmt.Sprintf("%g is out of the int32 range", x))
}
