// Exercise 6.1
// Write a function which accepts 2 integers and returns their sum, product and difference. Make a
// version with unnamed return variables, and a 2nd version with named return variables.

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。

package main

import "fmt"

var num1, num2 int = 5, 2

func main() {
    num_s, num_p, num_d := calcSum(num1, num2)
    fmt.Printf("Sum/Product/Diff: %d, %d: %d\n", num_s, num_p, num_d)
    num_s, num_p, num_d = calcSum_2(num1, num2)
    fmt.Printf("Sum/Product/Diff: %d, %d, %d\n", num_s, num_p, num_d)
}

func calcSum(num1, num2 int) (int, int, int) {
    return num1 + num2, num1 * num2, num1 - num2
}

func calcSum_2(num1, num2 int) (num_s, num_p, num_d int) {
    num_s = num1 + num2
    num_p = num1 * num2
    num_d = num1 - num2
    return
}

/*
Sum/Product/Diff: 7, 10: 3
Sum/Product/Diff: 7, 10, 3
*/
