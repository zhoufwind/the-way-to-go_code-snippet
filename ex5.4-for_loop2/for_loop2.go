package main

import "fmt"

func main() {
    i := 0
START:
    fmt.Printf("The counter is at %d\n", i)
    i++
    if i < 15 {
        goto START
    }
}

/*
The counter is at 0
The counter is at 1
The counter is at 2
The counter is at 3
The counter is at 4
The counter is at 5
The counter is at 6
The counter is at 7
The counter is at 8
The counter is at 9
The counter is at 10
The counter is at 11
The counter is at 12
The counter is at 13
The counter is at 14
*/
