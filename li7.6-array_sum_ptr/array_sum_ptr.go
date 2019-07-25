// 将数组传递给函数 - 传递数组的指针

package main

import "fmt"

func main() {
    array := [3]float64{7.0, 8.5, 9.1}
    x := Sum(&array)
    fmt.Printf("The sum of the array(%v) is: %f\n", array, x)
}

func Sum(a *[3]float64) (sum float64) {
    for _, v := range a {
        sum += v
    }
    return
}

/*
The sum of the array([7 8.5 9.1]) is: 24.600000
*/
