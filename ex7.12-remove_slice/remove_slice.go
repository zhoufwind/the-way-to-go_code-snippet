// Exercise 7.12
// Make a function RemoveStringSlice that removes items from index start to end in
// a slice. 

// 写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片中移除。

package main

import "fmt"

func main() {
    sl := []int{1,2,3,4,5}
    fmt.Printf("slice(%v) length: %d, capacity: %d\n", sl, len(sl), cap(sl))

    start := 1
    end := 3
    slnew := RemoveStringSlice(sl, start, end)
    fmt.Printf("New slice(%v) length: %d, capacity: %d\n", slnew, len(slnew), cap(slnew))
}

func RemoveStringSlice(sl []int, start, end int) (slnew []int) {
    sl1 := sl[:start]
    sl2 := sl[end:]
    slnew = append(sl1, sl2...)
    return
}

/*
slice([1 2 3 4 5]) length: 5, capacity: 5
New slice([1 4 5]) length: 3, capacity: 5
*/
