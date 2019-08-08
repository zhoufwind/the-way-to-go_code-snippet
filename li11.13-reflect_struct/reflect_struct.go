package main

import (
    "fmt"
    "reflect"
)

type NotknownType struct {
    s1, s2, s3 string
}

func (n NotknownType) String() string {
    return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate
var secret interface {} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
    value := reflect.ValueOf(secret)

    typ := reflect.TypeOf(secret)
    // alternative
    //typ := value.Type()
    fmt.Println(typ)

    knd := value.Kind()
    fmt.Println(knd)

    // iterate through the fields of the struct
    for i := 0; i < value.NumField(); i++ {
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
    }

    // call the first method, which is String()
    results := value.Method(0).Call(nil)
    fmt.Println(results)
}

/*
main.NotknownType
struct
Field 0: Ada
Field 1: Go
Field 2: Oberon
[Ada - Go - Oberon]
*/
