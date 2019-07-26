// Exercise 10.2
// Make a version of personex1.go where the parameter of upPerson is not a pointer.
// Explain the difference in behavior.

// 修改 persionext1.go，使它的参数 upPerson 不是一个指针，解释下二者的区别。

package main

import (
    "fmt"
    "strings"
)

type Person struct {
    firstName	string
    lastName	string
}

func upPerson(p Person) {
    p.firstName = strings.ToUpper(p.firstName)
    p.lastName = strings.ToUpper(p.lastName)
}

func main() {
    pers := new(Person)
    pers.firstName = "Stephen"
    pers.lastName = "Zhou"
    upPerson(*pers)
    fmt.Printf("Value: %v, position: %p\n", *pers, pers)
}

/*
Value: {Stephen Zhou}, position: 0xc42000a060
*/
