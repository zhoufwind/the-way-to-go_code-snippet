package main

import "fmt"

func main() {
    m := 7

    switch {
        case m >= 1 && m <= 3:
            fmt.Println("Q1")
        case m >= 4 && m <= 6:
            fmt.Println("Q2")
        case m >= 7 && m <= 9:
            fmt.Println("Q3")
        case m >= 10 && m <= 12:
            fmt.Println("Q4")
        default:
            fmt.Println("Invalid month!")
    }
}

/*
Q3
*/
