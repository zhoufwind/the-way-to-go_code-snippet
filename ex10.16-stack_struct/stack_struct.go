// Exercise 10.16
// Implement the stack datastructure:
// ----
// | l |
// ----
// | k |
// ----
// | j |
// ----
// | i |
// ----
// It has cells to contain data, for example integers i, j, k, l, etc. the cells are indexed
// from the bottom (index 0) to the top (index n). Let’s assume n=3 for this exercise,
// so we have 4 places.
// A new stack contains 0 in all cells.
// A new value is put in the highest cell which is empty (contains 0), on top: this is
// called push.
// To get a value from the stack, take the highest value which is not 0: this is called pop.
// So we can understand why a stack is called a Last In First Out (LIFO) structure.
// Define a new type Stack for this datastructure. Make 2 methods Push and Pop.
// Make a String() method (for debugging purposes) which shows the content of the
// stack as: [0:i] [1:j] [2:k] [3:l]
// (1) take as underlying data structure an array of 4 ints: stack_arr.go
// (2) take as underlying data structure a struct containing an index and an array of
// int, the index contains the first free position: stack_struct.go
// (3) generalize both implementations by making the number of elements 4 a
// constant LIMIT.

// 实现栈（stack）数据结构：
// 它的格子包含数据，比如整数 i、j、k 和 l 等等，格子从底部（索引 0）至顶部（索引 n）来索引。这个例子中假定 n=3，那么一共有 4 个格子。
// 一个新栈中所有格子的值都是 0。
// 将一个新值放到栈的最顶部一个空（包括零）的格子中，这叫做push。
// 获取栈的最顶部一个非空（非零）的格子的值，这叫做pop。 现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧。
// 为栈定义一个Stack 类型，并为它定义 Push 和 Pop 方法，再为它定义 String() 方法（用于调试）输出栈的内容，比如：[0:i] [1:j] [2:k] [3:l]。
// 1）stack_arr.go：使用长度为 4 的 int 数组作为底层数据结构。
// 2）stack_struct.go：使用包含一个索引和一个 int 数组的结构体作为底层数据结构，索引表示第一个空闲的位置。
// 3）使用常量 LIMIT 代替上面表示元素个数的 4 重新实现上面的 1）和 2），使它们更具有一般性。

package main

import (
    "fmt"
    "strconv"
)

const LIMIT = 4

type Stack struct {
    ix	int	// first free position, so data[ix] == 0
    data	[LIMIT]int
}

func (st *Stack) Push(n int) {
    if st.ix + 1 > LIMIT {
        return	// stack is full
    }
    st.data[st.ix] = n
    st.ix++
}

func (st *Stack) Pop() int {
    if st.ix == 0 {
        return 0	// stack is empty
    } else {
        st.ix--
        return st.data[st.ix]
    }
}

func (st *Stack) String() string {
    str := ""
    for ix := 0; ix < st.ix; ix++ {
        str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "]"
    }
    return str
}

func main() {
    st1 := new(Stack)
    fmt.Printf("%v\n", st1)
    fmt.Println()

    st1.Push(3)
    fmt.Printf("%v\n", st1)
    st1.Push(7)
    fmt.Printf("%v\n", st1)
    st1.Push(10)
    fmt.Printf("%v\n", st1)
    st1.Push(99)
    fmt.Printf("%v\n", st1)
    fmt.Println()
    //st1.Push(100)
    //fmt.Printf("%v\n", st1)
    //fmt.Println()

    p := st1.Pop()
    fmt.Printf("Popped: %d\n", p)
    fmt.Printf("%v\n", st1)
    p = st1.Pop()
    fmt.Printf("Popped: %d\n", p)
    fmt.Printf("%v\n", st1)
    p = st1.Pop()
    fmt.Printf("Popped: %d\n", p)
    fmt.Printf("%v\n", st1)
    p = st1.Pop()
    fmt.Printf("Popped: %d\n", p)
    fmt.Printf("%v\n", st1)
    //p = st1.Pop()
    //fmt.Printf("Popped: %d\n", p)
    //fmt.Printf("%v\n", st1)
}

/*


[0:3]
[0:3][1:7]
[0:3][1:7][2:10]
[0:3][1:7][2:10][3:99]

Popped: 99
[0:3][1:7][2:10]
Popped: 10
[0:3][1:7]
Popped: 7
[0:3]
Popped: 3

*/
