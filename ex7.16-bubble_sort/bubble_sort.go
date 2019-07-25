// Exercise 7.16
// Sort a slice of ints through a function which implements the Bubblesort-algorithm
// (look up the definition of this algorithm at http://en.wikipedia.org/wiki/
// Bubble_sort)

// 编写一个程序，使用冒泡排序的方法排序一个包含整数的切片（算法的定义可参考 维基百科）。

package main

import "fmt"

func main() {
    a := []int{4, 3, 2, 1}
    fmt.Printf("%v\n", a)
    b := BubbleSort(a)
    fmt.Printf("%v\n", b)
}

func BubbleSort(a []int) []int {
    for i := 0; i < len(a)-1; i++ {
        for j := 0; j < len(a)-1; j++ {
            if a[j] > a[j+1] {
                tmp := a[j+1]
                a[j+1] = a[j]
                a[j] = tmp
            }
        }
    }
    return a
}

/*
[4 3 2 1]
[1 2 3 4]
*/
