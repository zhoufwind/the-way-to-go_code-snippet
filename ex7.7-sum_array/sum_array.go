// Exercise 7.7
// Write a function Sum which has as parameter an array arrF of 4 floating-point
// numbers, and which returns the sum of all the numbers in the array sum_array.go
// How would the code have to be modified to use a slice instead of an array ?
// The slice-form of the function works for arrays of different lengths!

//写一个 Sum 函数，传入参数为一个 4 位 float 数组成的数组 arrF，返回该数组的所有数字和。
//如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？

package main

import "fmt"

func main() {
    arr1 := [4]float64{1.1, 2.2, 3.3, 4.4}
    fmt.Println(Sum(&arr1))
}

func Sum(a *[4]float64) (sum float64) {
    for _, v := range a {
        sum += v
    }
    return
}
