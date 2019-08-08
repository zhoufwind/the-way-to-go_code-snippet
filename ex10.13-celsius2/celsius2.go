// Exercise 10.13
// Make an alias type Celsius for float64 and define a String() method for it which
// prints out the temperature with 1 decimal and °C.

// 为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。

package main

import (
    "fmt"
    "strconv"
)

type Celsius float64

func (c Celsius) String() string {
    return "The temperature is: " + strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"
}

func main() {
    var c Celsius = 28.8
    fmt.Println(c)
}

/*
The temperature is: 28.8 °C
*/
