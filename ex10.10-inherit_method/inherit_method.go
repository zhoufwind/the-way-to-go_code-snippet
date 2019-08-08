// Exercise 10.10
// Define a struct type Base which contains an id-field, and methods Id() to return
// the id and SetId() to change the id. The struct type Person contains Base, and
// fields FirstName and LastName. The struct type Employee contains a Person and
// a salary.
// Make an employee instance and show its id.

// 定义一个结构体类型 Base，它包含一个字段 id，方法 Id() 返回 id，方法 SetId() 修
// 改 id。结构体类型 Person包含 Base，及 FirstName 和 LastName 字段。结构体类型 
// Employee 包含一个 Person 和 salary 字段。
// 创建一个 employee 实例，然后显示它的 id。

package main

import "fmt"

type Base struct {
    id	int
}

func (b *Base) Id() int {
    return b.id
}

func (b *Base) SetId(id int) {
    b.id = id
}

type Person struct {
    Base
    FirstName, LastName	string
}

type Employee struct {
    Person
    salary	float64
}

func main() {
    e := new(Employee)
    e.SetId(5)
    fmt.Println(e.Id())
    fmt.Println(e)
}

/*
5
&{{{5}  } 0}
*/
