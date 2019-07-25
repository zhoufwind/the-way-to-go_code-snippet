package main

import "fmt"

func f(a [3]int) { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }

func main() {
    var ar = [3]int{1, 2, 3}
    f(ar)	// pass a copy of ar
    fp(&ar)	// pass a pointer to ar
}

/*
[1 2 3]
&[1 2 3]
*/
