package main

import "fmt"

type struct1 struct {
    i1 int
    f1 float32
    str string
}

func main() {
    var ms *struct1 = new(struct1)
    // ms := new(struct1)
    ms.i1 = 10
    ms.f1 = 15.5
    ms.str = "Chris"

    fmt.Printf("The int is: %d\n", ms.i1)
    fmt.Printf("The float is: %f\n", ms.f1)
    fmt.Printf("The string is: %s\n", ms.str)
    fmt.Printf("The struct: %v\n", ms)
}

/*
The int is: 10
The float is: 15.500000
The string is: Chris
The struct: &{10 15.5 Chris}
*/
