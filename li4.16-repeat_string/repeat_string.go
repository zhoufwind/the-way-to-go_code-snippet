package main

import "fmt"
import "strings"

func main() {
    var oriS string = "Hi there! "
    var newS string

    newS = strings.Repeat(oriS, 3)
    fmt.Printf("The new repeated string is: %s\n", newS)
}

/*
The new repeated string is: Hi there! Hi there! Hi there! 
*/
