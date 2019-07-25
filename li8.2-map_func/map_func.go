// To demonstrate that the value can be any type, here is an example of a map which has a func() int
// as its value

package main

import "fmt"

func main() {
    mf := map[int]func() int {
        1: func() int { return 10 },
        2: func() int { return 20 },
        3: func() int { return 50 },
    }
    fmt.Println(mf)
}

/*
map[1:0x10935a0 2:0x10935b0 3:0x10935c0]
*/
