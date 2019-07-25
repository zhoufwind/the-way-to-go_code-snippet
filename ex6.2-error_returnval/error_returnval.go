// Exercise 6.2
// Write a function MySqrt which calculates the square root of a float64, but returns an error if this
// number is negative; make a version with unnamed and a second one with named return variables.

// 编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。

package main

import "fmt"
import "errors"
import "math"

var num1, num2 float64 = 25, -9

func main() {
    fmt.Printf("First example with %f: ", num1)
    ret1, err1 := MySqrt(num1)
    if err1 != nil {
        fmt.Println("Error! Return values are: ", ret1, err1)
    } else {
        fmt.Println("It's ok! Return values are: ", ret1, err1)
    }

    fmt.Printf("Second example with %f: ", num2)
    ret2, err2 := MySqrt2(num2)
    if err2 != nil {
        fmt.Println("Error! Return values are: ", ret2, err2)
    } else {
        fmt.Println("It's ok! Return values are: ", ret2, err2)
    }
}

func MySqrt(f float64) (float64, error) {
    if f < 0 {
        return float64(math.NaN()), errors.New("Negative number!")
    }
    return math.Sqrt(f), nil
}

func MySqrt2(f float64) (ret float64, err error) {
    if f < 0 {
        ret = float64(math.NaN())
        err = errors.New("Negative number!")
    } else {
        ret = math.Sqrt(f)
    }
    return
}

/*
First example with 25.000000: It's ok! Return values are:  5 <nil>
Second example with -9.000000: Error! Return values are:  NaN Negative number!
*/
