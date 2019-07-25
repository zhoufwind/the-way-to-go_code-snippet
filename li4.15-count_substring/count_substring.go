package main

import "fmt"
import "strings"

func main() {
    var str string = "Hello, how is it going, Hugo?"
    var manyG = "gggggggggg"

    fmt.Printf("Number of H's in %s is: ", str)
    fmt.Printf("%d\n", strings.Count(str, "H"))

    fmt.Printf("Number of double g's in %s is: %d\n", manyG, strings.Count(manyG, "gg"))
}

/*
Number of H's in Hello, how is it going, Hugo? is: 2
Number of double g's in gggggggggg is: 5
*/
