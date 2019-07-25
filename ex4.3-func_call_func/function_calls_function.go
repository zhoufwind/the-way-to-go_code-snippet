// Deduce the output of the following programs and explain your answer, then compile
// and execute them.

package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "0"
	print(a)
	f2()
}

func f2() {
	print(a)
}

// f1中定义局部变量，不影响全局变量值，f2中读取全局变量值
// G0G
