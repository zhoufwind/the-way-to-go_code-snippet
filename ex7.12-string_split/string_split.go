// Exercise 7.12
// Write a function Split with parameters a string to split and the position to split,
// which returns the two substrings.

// 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引i，然后返回两个分割后的字符串。

package main

import "fmt"

func main() {
    a := "hello"
    i := 2
    var b1, b2 = stringSplit(a, i)
    fmt.Printf("%s, %s\n", b1, b2)
}

func stringSplit(a string, i int) (b1 string, b2 string) {
    sl := []byte(a)
    sl1 := sl[0:i]
    sl2 := sl[i:]

    b1 = string(sl1)
    b2 = string(sl2)

    return
}

/*
he, llo
*/
