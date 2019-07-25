// Exercise 7.6
// Split a buffer buf into 2 slices: the header is the first n bytes, the tail is the rest; use
// 1 line of code.

// 把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。

package main

import "fmt"

func main() {
    sl := []byte{'a', 'b', 'c', 'd'}
    n := 1
    sl1 := sl[0:n]
    fmt.Printf("%c\n", sl1)
    sl2 := sl[n:]
    fmt.Printf("%c\n", sl2)
}
