// Exercise 11.12
// Make a slight variation to Ex. 11.11 to allow for mapFunc to receive a variable number of items.

// 稍微改变练习 11.11，允许 mapFunc 接收不定数量的 items。

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

    res1 := mapFunc(mf, 0, 1, 2, 3, 4, 5)
    for _, v := range res1 {
        fmt.Println(v)
    }
    println()

    res2 := mapFunc(mf, "0", "1", "2", "3", "4", "5")
    for _, v := range res2 {
        fmt.Println(v)
    }
}

func mapFunc(mf func(obj) obj, list ...obj) []obj {
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
