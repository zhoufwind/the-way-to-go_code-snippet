// Exercise 10.6
// Define a struct employee with a field salary, and make a method giveRaise for this type to increase
// the salary with a certain percentage.

// 定义结构体 employee，它有一个 salary 字段，给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水。

package main

import "fmt"

type employee struct {
    salary	float64
}

func main() {
    e1 := new(employee)
    e1.salary = 55000.0
    var raisePerc float64 = 10.0
    fmt.Printf("New salary: %f\n", e1.giveRaise(raisePerc))
}

func (e *employee) giveRaise(perc float64) (res float64) {
    res = e.salary * (1 + perc / 100)
    return
}

/*
New salary: 27500.000000
*/
