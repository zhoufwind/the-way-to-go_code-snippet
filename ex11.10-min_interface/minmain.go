// Exercise 11.10
// Analogous to the Sorter interface we developed in § 11.7, make a Miner interface with the necessary
// operations, and a function Min which has as a parameter a variable which is a collection of type
// Miner and which calculates and returns the minimum element in that collection.

// 仿照11.7中开发的 Sorter 接口，创建一个 Miner 接口并实现一些必要的操作。函数 Min 接受一个 Miner 类
// 型变量的集合，然后计算并返回集合中最小的元素。

package main

import (
    "fmt"
    "./min"
)

func ints() {
    data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
    a := min.IntArray(data)	// conversion to type IntArray
    m := min.Min(a)
    fmt.Printf("The minimum of the array is: %d\n", m)
}

func strings() {
    data := []string{"stephen", "stan", "cuan", "zac"}
    a := min.StringArray(data)
    m := min.Min(a)
    fmt.Printf("The minimum of the array is: %s\n", m)
}

func main() {
    ints()
    strings()
}

/*
The minimum of the array is: -5467984
The minimum of the array is: cuan
*/
