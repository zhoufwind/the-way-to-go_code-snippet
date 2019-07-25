package main

import "fmt"

func main() {
    a := []int{4, 3, 2, 1}
    fmt.Printf("%v\n", a)
    b := SelectionSort(a)
    fmt.Printf("%v\n", b)
}

func SelectionSort(a []int) []int {
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if a[j] < a[i] {
                tmp := a[j]
                a[j] = a[i]
                a[i] = tmp
            }
        }
    }
    return a
}

/*
[4 3 2 1]
[1 2 3 4]
*/
