// Exercise 9.1
// Use the package container/list to construct a double linked list, put the values 101,102,103 in it
// and make a printout of the list.

// 使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。

package main

import "fmt"
import "container/list"

func main() {
    l := list.New()
    l.PushBack(100)
    l.PushBack(101)
    l.PushBack(102)
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}

/*
100
101
102
*/
