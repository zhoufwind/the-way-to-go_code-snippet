// Exercise 11.11
// In Ex. 7.13 we defined a map function to apply to a slice of ints (map_function.go).
// With the help of the empty interface and the type switch we can now write a generic map function,
// which can be applied to many types. Construct a map-function mapFunc for ints and strings,
// which doubles the ints and concatenates the string with itself.
// Hint: For readability define an alias for interface { }, like: type obj interface{}

// 在练习 7.13 中我们定义了一个 map 函数来使用 int 切片 （map_function.go）。
// 通过空接口和类型断言，现在我们可以写一个可以应用于许多类型的 泛型 的 map 函数，为 int 和 string 构
// 建一个把 int 值加倍和将字符串值与其自身连接（译者注：即"abc"变成"abcabc"）的 map 函数 mapFunc。
// 提示：为了可读性可以定义一个 interface{} 的别名，比如：type obj interface{}

package main

import "fmt"

type obj interface{}

func main() {
    // define a generic lambda function mf
    mf := func(i obj) obj {
        switch i.(type) {
        case int:
            return i.(int) * 2
        case string:
            return i.(string) + i.(string)
        }
        return i
    }

    isl := []obj{0, 1, 2, 3, 4, 5}
    res1 := mapFunc(mf, isl)
    for _, v := range res1 {
        fmt.Println(v)
    }
    println()

    ssl := []obj{"0", "1", "2", "3", "4", "5"}
    res2 := mapFunc(mf, ssl)
    for _, v := range res2 {
        fmt.Println(v)
    }
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
    result := make([]obj, len(list))
    for ix, v := range list {
        result[ix] = mf(v)
    }
    return result
}

/*
0
2
4
6
8
10

00
11
22
33
44
55
*/
