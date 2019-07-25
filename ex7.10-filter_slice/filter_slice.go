// Exercise 7.10
// Using a higher order function for filtering a collection:
// s is a slice of the first 10 integers. Make a function Filter which accepts s as 1st
// parameter and a fn func(int) bool as 2nd parameter and returns the slice of the
// elements of s which fulfil the function fn (make it true). Test this out with fn
// testing if the integer is even.

// 用顺序函数过滤容器：s 是前 10 个整形的切片。构造一个函数 Filter，第一个参数是 s，
// 第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法
// 测试当整型值是偶数时的情况。

package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	input := 11

	res := filter(slice, matchEven)
	res2 := checkNum(slice, input)

	fmt.Println(res)
	fmt.Printf("Check Number res for %d is %t\n", input, res2)
}

func filter(s []int, fn func(int) bool) (res []int) {
	for ix, _ := range s {
		flag := fn(ix)
		if flag {
			res = append(res, s[ix])
		}
	}
	return
}

func matchEven(a int) (even bool) {
	if a%2 == 0 {
		even = true
	} else {
		even = false
	}
	return
}

func checkNum(s []int, input int) (nflag bool) {
	for i := 0; i < len(s); i++ {
		if s[i] == input {
			nflag = true
			break
		}
	}
	return
}

/*
[0 2 4 6 8 10 12]
Check Number res for 11 is true
*/
