// Question 7.2
// Given the slice of bytes b := []byte{‘g’, ‘o’, ‘l’, ‘a’, ‘n’, ‘g’}
// What are: b[1:4], b[:2], b[2:] and b[:] ?

// 给定切片 b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}，那么 b[1:4]、b[:2]、b[2:] 和 b[:] 分别是什么？

package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

	fmt.Printf("%c\n", b[1:4]) // o l a
	fmt.Printf("%c\n", b[:2])  // g o
	fmt.Printf("%c\n", b[2:])  // l a n g
	fmt.Printf("%c\n", b[:])   // g o l a n g
}

/*
[o l a]
[g o]
[l a n g]
[g o l a n g]
*/
