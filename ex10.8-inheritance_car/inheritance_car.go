// Exercise 10.8
// Make a working example for the Car and Engine code above; also give the Car type a field
// wheelCount and a method numberOfWheels() to retrieve this.
// Make a type Mercedes which embeds Car, an object of type Mercedes and use the methods.
// Then construct a method sayHiToMerkel() only on type Mercedes and invoke it.

// 创建一个上面 Car 和 Engine 可运行的例子，并且给 Car 类型一个 wheelCount 字段和一个 numberOfWheels() 方法。
// 创建一个 Mercedes 类型，它内嵌 Car，并新建 Mercedes 的一个实例，然后调用它的方法。
// 然后仅在 Mercedes 类型上创建方法 sayHiToMerkel() 并调用它。

package main

import "fmt"

type Engine interface {
    Start()
    Stop()
}

type Car struct {
    wheelCount	int
    Engine
}

func (c *Car) numberOfWheels() int {
    return c.wheelCount
}

func (c *Car) Start() {
    fmt.Println("Car is stared!")
}

func (c *Car) Stop() {
    fmt.Println("Car is stopped!")
}

func (c *Car) GoToWorkIn() {
    // get in car
    c.Start()
    // drive to work
    c.Stop()
    // get out of car
}

type Mercedes struct {
    Car
}

func (m *Mercedes) sayHiToMerkel() {
    fmt.Println("Hi Merkel!")
}

func main() {
    m := Mercedes{Car{4, nil}}
    fmt.Println(m.numberOfWheels())
    m.GoToWorkIn()
    m.sayHiToMerkel()
}

/*
4
Car is stared!
Car is stopped!
Hi Merkel!
*/
