package main

import "fmt"

func main() {
    slice1 := make([]int, 0, 10)

    // load the slice, cap(slice1) is 10:
    for i := 0; i < cap(slice1); i++ {
        slice1 = slice1[0:i+1]
        slice1[i] = i
        fmt.Printf("The length of slice is: %d\n", len(slice1))
    }

    // print the slice
    for ix, v := range slice1 {
        fmt.Printf("Slice at %d is: %d\n", ix, v)
    }

    slice2 := [10]int{0,1,2,3,4,5,6,7,8,9}
    sl2 := slice2[5:7]
    fmt.Printf("slice2(%v) length: %d, capacity: %d\n", sl2, len(sl2), cap(sl2))	// 5 6, 2, 5 
    sl2 = sl2[0:4]
    fmt.Printf("slice2(%v) length: %d, capacity: %d\n", sl2, len(sl2), cap(sl2))  // 5 6 7 8, 4, 5
}

/*
The length of slice is: 1
The length of slice is: 2
The length of slice is: 3
The length of slice is: 4
The length of slice is: 5
The length of slice is: 6
The length of slice is: 7
The length of slice is: 8
The length of slice is: 9
The length of slice is: 10
Slice at 0 is: 0
Slice at 1 is: 1
Slice at 2 is: 2
Slice at 3 is: 3
Slice at 4 is: 4
Slice at 5 is: 5
Slice at 6 is: 6
Slice at 7 is: 7
Slice at 8 is: 8
Slice at 9 is: 9
slice2([5 6]) length: 2, capacity: 5
slice2([5 6 7 8]) length: 4, capacity: 5
*/
