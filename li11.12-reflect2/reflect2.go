package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))

    v := reflect.ValueOf(x)
    // setting a value:
    // Error: will panic: reflect: reflect.Value.SetFloat using unaddressable value
    // v.SetFloat(3.1415)

    fmt.Println("settability of v:", v.CanSet())
    v = reflect.ValueOf(&x)	// Note: take the address of x
    fmt.Println("type of v:", v.Type())
    fmt.Println("settability of v:", v.CanSet())
    v = v.Elem()
    fmt.Println("The Elem of v is:", v)
    fmt.Println("settability of v:", v.CanSet())
    v.SetFloat(3.1415)	// this works

    fmt.Println("value:", v)

    fmt.Println("type:", v.Type())
    fmt.Println("kind:", v.Kind())
    fmt.Println("value:", v.Float())

    fmt.Println(v.Interface())
    fmt.Printf("value is %5.2e\n", v.Interface())

    y := v.Interface().(float64)
    fmt.Println(y)
}

/*
type: float64
settability of v: false
type of v: *float64
settability of v: false
The Elem of v is: 3.4
settability of v: true
value: 3.1415
type: float64
kind: float64
value: 3.1415
3.1415
value is 3.14e+00
3.1415
*/
