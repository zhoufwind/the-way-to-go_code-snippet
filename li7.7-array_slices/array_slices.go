package main

import "fmt"

func main() {
    var arr1 [6]int
    var slice1 []int = arr1[2:5]	// 从2开始包含2，不包含5，也即：2,3,4

    // load the array with integers: 0, 1, 2, 3, 4, 5
    for i := 0; i < len(arr1); i++ {
        arr1[i] = i
    }

    // Print the whole slice1
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }

    fmt.Printf("The length of arr1 is: %d\n", len(arr1))
    fmt.Printf("The length of slice1 is: %d\n", len(slice1))
    // 如果 s 是一个切片，cap(s) 就是从 s[0] 到数组末尾的数组长度。
    fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

    // grow the slice
    slice1 = slice1[0:4]
    for i := 0; i < len(slice1); i++ {
        fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    }

    fmt.Printf("The length of slice1 is: %d\n", len(slice1))
    fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

/*
Slice at 0 is 2
Slice at 1 is 3
Slice at 2 is 4
The length of arr1 is: 6
The length of slice1 is: 3
The capacity of slice1 is 4
Slice at 0 is 2
Slice at 1 is 3
Slice at 2 is 4
Slice at 3 is 5
The length of slice1 is: 4
The capacity of slice1 is 4
*/
