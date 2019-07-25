// Exercise 7.2
// Write the loop that fills an array with the loop-counter (from 0 to
// 15) and then prints that array to the screen.

// 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。

package main

import "fmt"

func main() {
	var arr1 [16]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for _, v := range arr1 {
		fmt.Println(v)
	}
}

/*
0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
*/
