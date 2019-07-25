// Exercise 7.1
// Prove that when assigning an array to another, a distinct copy in
// memory of the array is made.

// 证明当数组赋值时，发生了数组内存拷贝。

package main

import "fmt"

func main() {
    var arr1 [5]int

    for i := 0; i < len(arr1); i++ {
        arr1[i] = i * 2
    }

    fmt.Printf("arr1's position: %p\n", &arr1)
    for i := 0; i < len(arr1); i++ {
        fmt.Printf("arr1 index: %d, value: %d, position: %p\n", i, arr1[i], &arr1[i])
    }
    fmt.Println()

    arr2 := arr1
    arr2[2] = 100

    fmt.Printf("arr1's position: %p\n", &arr1)
    for i := 0; i < len(arr1); i++ {
        fmt.Printf("arr1 index: %d, value: %d, position: %p\n", i, arr1[i], &arr1[i])
    } 
    fmt.Println()

    fmt.Printf("arr2's position: %p\n", &arr2)
    for i := 0; i < len(arr2); i++ {
        fmt.Printf("arr2 index: %d, value: %d, position: %p\n", i, arr2[i], &arr2[i])
    }
}

/*
arr1's position: 0xc42001a060
arr1 index: 0, value: 0, position: 0xc42001a060
arr1 index: 1, value: 2, position: 0xc42001a068
arr1 index: 2, value: 4, position: 0xc42001a070
arr1 index: 3, value: 6, position: 0xc42001a078
arr1 index: 4, value: 8, position: 0xc42001a080

arr1's position: 0xc42001a060
arr1 index: 0, value: 0, position: 0xc42001a060
arr1 index: 1, value: 2, position: 0xc42001a068
arr1 index: 2, value: 4, position: 0xc42001a070
arr1 index: 3, value: 6, position: 0xc42001a078
arr1 index: 4, value: 8, position: 0xc42001a080

arr2's position: 0xc42001a090
arr2 index: 0, value: 0, position: 0xc42001a090
arr2 index: 1, value: 2, position: 0xc42001a098
arr2 index: 2, value: 100, position: 0xc42001a0a0
arr2 index: 3, value: 6, position: 0xc42001a0a8
arr2 index: 4, value: 8, position: 0xc42001a0b0
*/
