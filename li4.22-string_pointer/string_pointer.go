package main

import "fmt"

func main() {
    s := "good bye"
    var p *string = &s
    fmt.Printf("Here is the pointer p: %p\n", p)	// Here is the pointer p: 0xc42000e1d0

    *p = "hello"
    fmt.Printf("Here is the string *p: %s\n", *p)	// Here is the string *p: hello

    fmt.Printf("Here is the string s: %s\n", s)		// Here is the string s: hello
}

/*
Here is the pointer p: 0xc42000e1d0
Here is the string *p: hello
Here is the string s: hello
*/
