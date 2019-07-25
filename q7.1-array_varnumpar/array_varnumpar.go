package main

import "fmt"

func main() {
    a := [...]string{"a", "b", "c", "d"}
    for i, _ := range a {
        fmt.Println("Array item", i, "is", a[i])
    }
}

/*
Array item 0 is a
Array item 1 is b
Array item 2 is c
Array item 3 is d
*/
