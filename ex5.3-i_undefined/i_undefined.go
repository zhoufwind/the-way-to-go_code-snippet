// Exercise 5.3
// this program does not compile, why not ?
// package main
// import “fmt”
// func main() {
// for i:=0; i<10; i++ {
//  fmt.Printf(“%v\n”, i)
// }
// fmt.Printf(“%v\n”, i) //<-- compile error: undefined i
// }
// How could you make it work ?

package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%v\n", i)
    }
    var i int
    fmt.Printf("%v\n", i)
}

/*
0
1
2
3
4
5
6
7
8
9
0
*/
