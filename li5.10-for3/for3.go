package main

import "fmt"

func main() {
    i := 10
    for {
        i = i - 1
        fmt.Printf("%d\n", i)
        if i < 0 {
            break
        }
    }
}

/*
9
8
7
6
5
4
3
2
1
0
-1
*/
