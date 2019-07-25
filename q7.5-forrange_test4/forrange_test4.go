package main

import "fmt"

func main() {
    a := []int{10, 20, 30, 40, 50}
    pa := &a
    fmt.Println(*pa)

    for i := 0; i < len(*pa); i++ {
        (*pa)[i] *= 2
    }

    fmt.Println(*pa)
    fmt.Println(a)
}

/*
[10 20 30 40 50]
[20 40 60 80 100]
[20 40 60 80 100]
*/
