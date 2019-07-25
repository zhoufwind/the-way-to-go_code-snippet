package main

import "fmt"

func main() {
    i := 0
    for {	// since there are no checks, this is an infinite loop
        if i >= 3 {	// break out of this for loop when this condition is met
            break
        }
        fmt.Println("Value of i is: ", i)
        i++
    }
    fmt.Println("A statement just after for loop.")
}

/*
Value of i is:  0
Value of i is:  1
Value of i is:  2
A statement just after for loop.
*/
