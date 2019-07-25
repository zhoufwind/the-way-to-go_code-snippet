package main

import (
    "fmt"
    "./fibo"
)

var nextFibo int
var op string

func main() {
    op = "+"
    calls()
    fmt.Println("Change of opertation from + to *")
    nextFibo = 0
    op = "*"
    calls()
}

func calls() {
    next()
    fmt.Println("...")
    next()
    fmt.Println("...")
    next()
    fmt.Println("...")
    next()
}

func next() {
    result := 0
    nextFibo++
    result = fibo.Fibonacci(op, nextFibo)
    fmt.Printf("fibonacci(%d) is: %d\n", nextFibo, result)
}

/*
fibonacci(1) is: 1
...
fibonacci(2) is: 2
...
fibonacci(3) is: 3
...
fibonacci(4) is: 5
Change of opertation from + to *
fibonacci(1) is: 2
...
fibonacci(2) is: 4
...
fibonacci(3) is: 8
...
fibonacci(4) is: 32
*/
