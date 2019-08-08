// Exercise 11.8
// Define a struct Person with firstName and LastName, and a type Persons as a [ ]Person.
// Implement the Sorter interface for Persons and test it.

// 定义一个结构体 Person，它有两个字段：firstName 和 lastName，为 []Person 定义类型 Persons 。
// 让 Persons实现 Sorter 接口并进行测试。

package main

import (
    "fmt"
    "./sort"
)

type Person struct {
    firstName, lastName	string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
    in := p[i].lastName + " " + p[i].firstName
    jn := p[j].lastName + " " + p[j].firstName
    return in < jn
}

func (p Persons) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func main() {
    p1 := Person{"Stephen", "Zhou"}
    p2 := Person{"Stan", "Qian"}
    p3 := Person{"Zac", "Gu"}
    p4 := Person{"Cuan", "Qu"}
    arrP := Persons{p1, p2, p3, p4}
    fmt.Printf("Before sorting: %v\n", arrP)
    sort.Sort(arrP)
    fmt.Printf("After sorting: %v\n", arrP)
}

/*
Before sorting: [{Stephen Zhou} {Stan Qian} {Zac Gu} {Cuan Qu}]
After sorting: [{Zac Gu} {Stan Qian} {Cuan Qu} {Stephen Zhou}]
*/
