// Deduce the output of the following programs and explain your answer, then compile
// and execute them.

package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a := "0"
	print(a)
}

// 局部变量，定义只局限于函数内部
// G0G
