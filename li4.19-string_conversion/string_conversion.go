package main

import "fmt"
import "strconv"

func main() {
    var ori string = "666"

    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

    an, _ := strconv.Atoi(ori)
    fmt.Printf("The integer is: %d\n", an)

    an = an + 5
    newS := strconv.Itoa(an)
    fmt.Printf("The new string is: %s\n", newS)
}

/*
The size of ints is: 64
The integer is: 666
The new string is: 671
*/
