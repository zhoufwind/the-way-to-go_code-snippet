package main

import "io"
import "log"

func func1(s string) (n int, err error) {
    defer func() {
        log.Printf("func1(%q) = %d, %v", s, n, err)
    }()
    return 7, io.EOF
}

func main() {
    func1("Go")
}

/*
2019/07/21 12:23:06 func1("Go") = 7, EOF
*/
