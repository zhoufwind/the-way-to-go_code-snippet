package main

import "fmt"

func main() {
    a := 1
    goto TARGET
    b := 9
        TARGET:
    b += a
    fmt.Printf("“a is %v *** b is %v", a, b)
}

/*
goto TARGET jumps over declaration of b at ./goto2.go:8:7
*/
