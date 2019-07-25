// Question 7.3
// Given s := make([]byte, 5), what is len(s) and cap(s) ?
// s = s[2:4], what is now len(s) and cap(s) ?

// 给定 s := make([]byte, 5)，len(s) 和 cap(s) 分别是多少？s = s[2:4]，len(s) 和 cap(s) 又分别是多少？

package main

import "fmt"

func main() {
	s := make([]byte, 5)
	fmt.Printf("length: %d, capacity: %d\n", len(s), cap(s)) // 5, 5

	s = s[2:4]
	fmt.Printf("length: %d, capacity: %d\n", len(s), cap(s)) // 2, 3
}

/*
length: 5, capacity: 5
length: 2, capacity: 3
*/
