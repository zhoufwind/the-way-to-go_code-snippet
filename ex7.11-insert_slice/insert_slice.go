// Exercise 7.11
// Make a function InsertStringSlice that inserts a slice into another slice at a certain
// index. 

// 写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。

package main

import "fmt"

func main() {
    sl1 := []int{10,20,30}
    fmt.Printf("sl1: %v\n", sl1)

    ix := 2
    fmt.Printf("index: %d, insert into sl1 between: %d and %d\n", ix, sl1[ix-1], sl1[ix])

    sl2 := []int{4,5,6}
    fmt.Printf("sl2: %v\n", sl2)

    sl3 := InsertStringSlice(sl1, ix, sl2)
    fmt.Printf("sl3: %v\n", sl3)
}

func InsertStringSlice(sl1 []int, ix int, sl2 []int) (sl3 []int) {
    sl1a := sl1[:ix]
    sl1b := sl1[ix:]
    sl3 = Append(sl1a, sl2)
    sl3 = Append(sl3, sl1b)
    return
}

func Append(slice []int, data []int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's need, for future growth
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

/*
sl1: [10 20 30]
index: 2, insert into sl1 between: 20 and 30
sl2: [4 5 6]
sl3: [10 20 4 5 6 30]
*/
