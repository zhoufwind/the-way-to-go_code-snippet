// Question 7.4
// Suppose s1 := []byte{‘p’, ‘o’, ‘e’, ‘m’} and s2 := d[2:]
// What is the value of s2 ?
// We do: s2[1] == ‘t’, what is now the value of s1 and s2 ?

// 假设 s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := s1[2:]，s2 的值是多少？如果我们执行 s2[1] = 't'，s1 和 s2 现在的值又分配是多少？

package main

import "fmt"

func main() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:] // e, m
	fmt.Printf("%c\n", s2)

	s2[1] = 't' // e, t
	fmt.Printf("%c\n", s2)
}

/*
[e m]
[e t]
*/
