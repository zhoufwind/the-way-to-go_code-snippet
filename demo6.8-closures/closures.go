package main

import "fmt"

func main() {
    sumNum := func(a, b int) int { return a + b }
    fmt.Printf("Sum: %d\n", sumNum(1, 2))
}

/*
Sum: 3
*/
