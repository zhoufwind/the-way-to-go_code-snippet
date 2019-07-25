package main

import "fmt"
import "math"

func main() {
    var a int = -1
    b, err := Uint8FromInt(a)
    fmt.Printf("%d, %v\n", b, err)
}

func Uint8FromInt(n int) (uint8, error) {
    if 0 <= n && n <= math.MaxUint8 {	// conversion is safe
        return uint8(n), nil
    }
    return 0, fmt.Errorf("%d is out of the uint8 range", n)
}
