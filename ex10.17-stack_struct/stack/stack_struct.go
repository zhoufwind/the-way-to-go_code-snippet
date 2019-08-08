// Exercise 10.17
// Starting from exercise 10.13 (the struct implementation of a stack datastructure),
// create a separate package stack for the stack implementation (stack_struct.go) and call it from a
// main program main_stack.go

// 从练习 10.16 开始（它基于结构体实现了一个栈结构），为栈的实现（stack_struct.go）创建一个单独的包 stack，并从 main 包 main.stack.go 中调用它。

package stack

import "strconv"

const LIMIT = 4

type Stack struct {
    ix	int
    data	[LIMIT]int
}

func (st *Stack) Push(n int) {
    if st.ix + 1 > LIMIT {
        return
    }
    st.data[st.ix] = n
    st.ix++
}

func (st *Stack) Pop() int {
    st.ix--
    return st.data[st.ix]
}

func (st *Stack) String() string {
    str := ""
    for ix := 0; ix < st.ix; ix++ {
        str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "]"
    }
    return str
}
