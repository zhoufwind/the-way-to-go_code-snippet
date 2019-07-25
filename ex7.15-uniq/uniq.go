// Exercise 7.15
// Write a program that traverses an array of characters, and copies them to another
// array only if the character is different from that which precedes it.

// 编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。

package main

import "fmt"

func main() {
    a := []byte{'b', 'b', 'c', 'd', 'd', 'd'}
    fmt.Printf("%c\n", a)
    b := uniq(a)
    fmt.Printf("%c\n", b)
}

func uniq(a []byte) (b []byte) {
    b = append(b, a[0])
    for i := 1; i < len(a); i++ {
        if a[i] != a[i - 1] {
            b = append(b, a[i])
        }
    }
    return b
}

/*
[b b c d d d]
[b c d]
*/
