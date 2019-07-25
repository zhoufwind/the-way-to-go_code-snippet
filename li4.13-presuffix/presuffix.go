package main

import "fmt"
import "strings"

func main() {
    var str string = "this is an pala string"
    fmt.Printf("T/F? Does the string \"%s\" have prefix \"%s\"?\n", str, "Th")
    fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

/*
T/F? Does the string "this is an pala string" have prefix "Th"?
false
*/
