package main

import "fmt"

func Sum(a []int) (sum int) {
    //sum = 0
    for i := 0; i < len(a); i++ {
        sum += i
    }
    return
}

func main() {
    var arr1 = [5]int{0, 1, 2, 3, 4}
    fmt.Println(Sum(arr1[:]))
}
