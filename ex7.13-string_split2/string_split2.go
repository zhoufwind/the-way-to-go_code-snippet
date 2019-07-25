// Exercise 7.13
// If str is a string, what then is str[len(str)/2:] + str[:len(str)/2]? Test it

// 假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？

package main

import "fmt"

func main() {
    str := "hello"
    fmt.Printf("string(%s) length: %d\n", str, len(str))
    fmt.Println(str[len(str)/2:])	// llo
    fmt.Println(str[:len(str)/2])	// he
    fmt.Println(str[len(str)/2:] + str[:len(str)/2])	// llohe
}

/*
string(hello) length: 5
llo
he
llohe
*/
