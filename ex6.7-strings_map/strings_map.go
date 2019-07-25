// Exercise 6.7
// The Map function in the package strings is also a good example of the use of higher order functions,
// like strings.IndexFunc(). Look up its definition in the package documentation and make a little
// test program with a map function that replaces all non-ASCII characters from a string with a ? or
// a space. What do you have to do to delete these characters?

// 包 strings 中的 Map 函数和 strings.IndexFunc() 一样都是非常好的使用例子。请学习它的源代码并基于该函
// 数书写一个程序，要求将指定文本内的所有非 ASCII 字符替换成 ? 或空格。您需要怎么做才能删除这些字符呢？

package main

import (
    "fmt"
    "strings"
)

func main() {
    asciiOnly := func(c rune) rune {
        if c > 127 {
            return ' '
        }
        return c
    }
    fmt.Println(strings.Map(asciiOnly, "München, 慕尼黑, Munich, ミュンヘン."))
}

/*
M nchen,    , Munich,      .
*/
