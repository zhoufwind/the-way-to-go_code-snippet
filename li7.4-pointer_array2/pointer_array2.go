package main

import "fmt"

// 取数组常量的地址来作为指向新实例的指针
func fp(a *[4]int) { fmt.Println(a) }

func main() {
    for i := 0; i < 5; i++ {
        fp(&[4]int{i, i*i, i*i*i, i*i*i*i})
    }
}

/*
&[0 0 0 0]
&[1 1 1 1]
&[2 4 8 16]
&[3 9 27 81]
&[4 16 64 256]
*/
