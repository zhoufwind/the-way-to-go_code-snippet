// Exercise 4.6
// Create a program that counts the number of bytes and characters (runes) for this string:
//  “asSASA ddd dsjkdsjs dk”
// Then do the same for this string: “asSASA ddd dsjkdsjsこん dk”
// Explain the difference. (hint: use the unicode/utf8 package.)

// 创建一个用于统计字节和字符（rune）的程序，并对字符串 asSASA ddd dsjkdsjs dk 进行分析，然后再分析 asSASA ddd dsjkdsjsこん dk，最后解释两者不同的原因（提示：使用 unicode/utf8 包）。

package main

import "fmt"
import "unicode/utf8"

func main() {
    // count number of characters
    str1 := "asSASA ddd dsjkdsjs dk"
    fmt.Printf("str1: '%s'\n", str1)
    fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
    fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))

    str2 := "asSASA ddd dsjkdsjsこん dk"
    fmt.Printf("str2: '%s'\n", str2)
    fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
    fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}

/*
str1: 'asSASA ddd dsjkdsjs dk'
The number of bytes in string str1 is 22
The number of characters in string str1 is 22
str2: 'asSASA ddd dsjkdsjsこん dk'
The number of bytes in string str2 is 28
The number of characters in string str2 is 24
*/
