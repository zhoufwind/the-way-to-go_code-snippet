package main

import "fmt"

func main() {
    map1 := make(map[int]float64)
    map1[1] = 1.0
    map1[2] = 2.0
    map1[3] = 3.0
    map1[4] = 4.0

    for key, value := range map1 {
        fmt.Printf("Key is: %d - value is: %f\n", key, value)
    }
}

/*
Key is: 1 - value is: 1.000000
Key is: 2 - value is: 2.000000
Key is: 3 - value is: 3.000000
Key is: 4 - value is: 4.000000
*/
