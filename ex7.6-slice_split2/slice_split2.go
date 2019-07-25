// Exercise 7.6
// Split a buffer buf into 2 slices: the header is the first n bytes, the tail is the rest; use
// 1 line of code.

// 把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。

package main

import "fmt"

func main() {
    str := "google"
    i := 3
    str1, str2 := splitByIndex(str, i)
    fmt.Printf("The string \"%s\" split at position %d is: %s / %s\n", str, i, str1, str2)
}

func splitByIndex(str string, i int) (str1, str2 string) {
    sl := []byte(str)
    sl1 := sl[:i]
    sl2 := sl[i:]
    str1 = string(sl1)
    str2 = string(sl2)
    return
}

/*
The string "google" split at position 3 is: goo / gle
*/
