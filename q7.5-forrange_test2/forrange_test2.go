package main
  
import "fmt"

func main() {
    arr1 := [5]int{10, 20, 30, 40, 50}
    fmt.Println(Double(arr1))
}

func Double(a [5]int) (d [5]int) {
    for ix, v := range a {
        d[ix] = v*2
    }
    return
}
