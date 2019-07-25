package main

import "fmt"

func main() {
    for i := 0; i < 7; i++ {
        if i % 2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}

/*
1
3
5
*/
