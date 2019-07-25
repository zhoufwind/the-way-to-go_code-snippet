package main

import "fmt"

func main() {
    var i1 = 5
    // An integer: 5, its location in memory: 0xc420014068
    fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

    var intP *int
    intP = &i1
    fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}

/*
An integer: 5, its location in memory: 0xc420014068
The value at memory location 0xc420014068 is 5
*/
