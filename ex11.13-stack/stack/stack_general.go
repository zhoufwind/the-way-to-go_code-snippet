package stack

import "errors"

type Stack []interface{}

// get stack's length
func (stack Stack) Len() int {
    return len(stack)
}

// get stack's capacity
func (stack Stack) Cap() int {
    return cap(stack)
}

// check if stack is empty
func (stack Stack) IsEmpty() bool {
    return len(stack) == 0
}

// push data
func (stack *Stack) Push(e interface{}) {
    *stack = append(*stack, e)
}

// get top data
func (stack Stack) Top() (interface{}, error) {
    if len(stack) == 0 {
        return nil, errors.New("stack is empty")
    }
    return stack[len(stack)-1], nil
}

// pop data
func (stack *Stack) Pop() (interface{}, error) {
    stk := *stack
    if len(stk) == 0 {
        return nil, errors.New("stack is empty")
    }
    top := stk[len(stk)-1]
    *stack = stk[:len(stk)-1]
    return top, nil
}
