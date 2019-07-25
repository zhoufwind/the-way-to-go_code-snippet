package main

import "fmt"

func main() {
    num1 := 7

    switch {
        case num1 < 0:
            fmt.Printf("Number is negative\n")
        case num1 > 0 && num1 < 10:
            fmt.Printf("Number is between 0 and 10\n")
        default:
            fmt.Printf("Number is 10 or greater\n")
    }
}

/*
Number is between 0 and 10
*/
