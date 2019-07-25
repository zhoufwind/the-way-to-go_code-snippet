package main

import "fmt"

func main() {
    f()
}

func f() {
    fv := func() { fmt.Println("Hello World") }
    fv()
    fmt.Printf("type: %T, value: %v\n", fv, fv)
}

/*
Hello World
type: func(), value: 0x1095390
*/
