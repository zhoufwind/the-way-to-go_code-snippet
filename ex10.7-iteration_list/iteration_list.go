// Exercise 10.7
// What’s wrong with the following code ?

package main

import "container/list"

func (p *list.List) Iter() {
    // ...
}

func main() {
    lst := new(list.List)
    for _ = range lst.Iter() {
        // ...
    }
}

/*
# command-line-arguments
./iteration_list.go:8:6: cannot define new methods on non-local type list.List
./iteration_list.go:14:22: lst.Iter undefined (type *list.List has no field or method Iter)
*/
