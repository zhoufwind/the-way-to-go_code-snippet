// Exercise 7.17
// In functional languages a map-function is a function which takes a function and a
// list as arguments, and its result is a list where the argument-function is applied to
// each element of the list, formally:
// map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), … F(en) )
// Write such a function mapFunc which takes as arguments:
// - a (lambda) function that multiplies an int by 10
// - a list of ints
// and returns the list of all ints multiplied by 10.

// 在函数式编程语言中，一个 map-function 是指能够接受一个函数原型和一个列表，并使用列表中的值依次执行函数原型，公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。
// 编写一个函数 mapFunc 要求接受以下 2 个参数：
// - 一个将整数乘以 10 的函数
// - 一个整数列表
// 最后返回保存运行结果的整数列表。

package main

import "fmt"

func main() {
    list := []int{0, 1, 2, 3, 4, 5, 6, 7}
    mf := func(i int) int {
        return i * 10
    }

    fmt.Printf("%v", mapFunc(mf, list))
    println()
}

func mapFunc(mf func(int) int, list []int) (result []int) {
    result = make([]int, len(list))
    for ix, v := range list {
        result[ix] = mf(v)
    }
    return
}

/*
[0 10 20 30 40 50 60 70]
*/
