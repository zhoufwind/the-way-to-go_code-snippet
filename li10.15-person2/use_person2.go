package main

import (
    "fmt"
    "./person2"
)

func main() {
    p := new(person2.Person)
    p.SetFirstName("Stephen")
    fmt.Println(p.FirstName())
}

/*
Stephen
*/
