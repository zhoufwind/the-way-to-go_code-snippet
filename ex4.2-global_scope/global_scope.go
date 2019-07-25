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
	a = "0"
	print(a)
}

// 全局变量，修改了全局变量值，影响全局
// G00
