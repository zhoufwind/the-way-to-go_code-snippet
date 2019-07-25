package main

import "fmt"

func main() {
    var arr1 [5]int

    for i := 0; i < len(arr1); i++ {
        arr1[i] = i * 2
    }

    for i := 0; i < len(arr1); i++ {
        fmt.Printf("Array at index %d is %d\n", i, arr1[i])
    }
}

/*
Array at index 0 is 0
Array at index 1 is 2
Array at index 2 is 4
Array at index 3 is 6
Array at index 4 is 8
*/
