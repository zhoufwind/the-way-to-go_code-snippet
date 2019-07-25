package main

import (
    "fmt"
    "./trans"
)

var twoPi = 2 * trans.Pi

func main() {
    fmt.Printf("2*Pi=%g\n", twoPi)
}

/*
2*Pi=6.283185307179586

tips:
%g -> 6.283185307179586
%f -> 6.283185
*/
