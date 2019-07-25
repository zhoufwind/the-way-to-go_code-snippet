// Exercise 7.5
// Given a slice sl we want to append a []byte data.
// Write a function Append(slice, data[]byte) []byte which lets sl grow if it is not
// big enough to accommodate data.

// 给定切片 sl，将一个 []byte 数组追加到 sl 后面。写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。

package main

import "fmt"

func main() {
	sl := []byte{'a', 'b', 'c'}
	fmt.Printf("sl(%c) length: %d, capacity: %d\n", sl, len(sl), cap(sl)) // 3, 3

	sl1 := []byte{'d', 'e'}
	fmt.Printf("sl1(%c) length: %d, capacity: %d\n", sl1, len(sl1), cap(sl1)) // 2, 2

	sl2 := Append(sl, sl1)
	fmt.Printf("sl2(%c) length: %d, capacity: %d\n", sl2, len(sl2), cap(sl2)) // m=3, n=3+2=5, 5, 12
}

func Append(slice []byte, data []byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's need, for future growth
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

/*
sl([a b c]) length: 3, capacity: 3
sl1([d e]) length: 2, capacity: 2
sl2([a b c d e]) length: 5, capacity: 12
*/
