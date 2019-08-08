package main

import "fmt"

type Node struct {
    le *Node
    data interface{}
    ri *Node
}

func NewNode(left, right *Node) *Node {
    return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
    n.data = data
}

func main() {
    root := NewNode(nil, nil)
    root.SetData("root node")

    // make child (leaf) nodes:
    a := NewNode(nil, nil)
    a.SetData("left node")
    b := NewNode(nil, nil)
    b.SetData("right node")

    root.le = a
    root.ri = b

    fmt.Printf("root: %v\n", root)
    fmt.Printf("left: %v, position: %p\n", a, &a)
    fmt.Printf("right: %v, position: %p\n", b, &b)
}

/*
root: &{0xc42000a080 root node 0xc42000a0a0}
left: &{<nil> left node <nil>}, position: 0xc42000c028
right: &{<nil> right node <nil>}, position: 0xc42000c030
*/
