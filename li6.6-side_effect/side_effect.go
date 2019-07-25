package main

import "fmt"

func main() {
	n := 0
	reply := &n
	fmt.Printf("reply's value: %d | reply's postion: %p | n's value: %d | n's postion: %p\n", *reply, reply, n, &n)
	Multiply(10, 5, reply)
	fmt.Printf("reply's value: %d | reply's postion: %p | n's value: %d | n's postion: %p\n", *reply, reply, n, &n)
}

// this function chages reply
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

/*
reply's value: 0 | reply's postion: 0xc420014068 | n's value: 0 | n's postion: 0xc420014068
reply's value: 50 | reply's postion: 0xc420014068 | n's value: 50 | n's postion: 0xc420014068
*/
