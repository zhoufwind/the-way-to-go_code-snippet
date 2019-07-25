package main

import "fmt"

func main() {
    Function1()
}

func Function1() {
    fmt.Printf("In Function1 at the top\n")
    defer Function2()
    fmt.Printf("In Function2 at the bottom\n")
}

func Function2() {
    fmt.Print("Function2: Deferred until the end of the calling function\n")
}

/*
In Function1 at the top
In Function2 at the bottom
Function2: Deferred until the end of the calling function
*/
