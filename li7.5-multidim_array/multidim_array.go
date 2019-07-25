package main

import "fmt"

const (
    WIDTH = 4
    HEIGHT = 3
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++ {
            screen[x][y] = 8
        }
    }

    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++ {
            fmt.Print(screen[x][y])
        }
        fmt.Println()
    }
}

/*
8888
8888
8888
*/
