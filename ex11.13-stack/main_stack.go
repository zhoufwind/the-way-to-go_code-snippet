// Exercise 11.13
// In exercises 10.10 and 10.11 we developed some Stack struct-types. However they were limited to
// a certain fixed internal type. Now develop a general stack type using a slice with as element type
// the interface{ }.
// Implement the following Stack-methods: Len() int, IsEmpty() bool, Push(x interface{}) and Pop()
// (x interface{}, error). Pop() changes the stack while returning the topmost element; write also a
// method Top() which only returns this element.
// In the main program construct a stack with a number of elements of different types, Pop all of
// them and print them.

// 在练习 10.16 和 10.17 中我们开发了一些栈结构类型。但是它们被限制为某种固定的内建类型。现在用一个元素类型是 interface{}（空接口）的切片开发一个通用的栈类型。
// 实现下面的栈方法：
// Len() int
// IsEmpty() bool
// Push(x interface{})
// Pop() (interface{}, error)
// Pop() 改变栈并返回最顶部的元素；Top() 只返回最顶部元素。
// 在主程序中构建一个充满不同类型元素的栈，然后弹出并打印所有元素的值。

package main

import (
    "fmt"
    "./stack"
)

var st1 stack.Stack

func main() {
    st1.Push("Stephen")
    st1.Push(3.14)
    st1.Push(100)
    st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})

    for {
        item, err := st1.Pop()
        if err != nil {
            break
        }
        fmt.Println(item)
    }
}

/*
[Java C++ Python C# Ruby]
100
3.14
Stephen
*/
