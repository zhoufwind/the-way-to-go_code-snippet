// Exercise 10.11
// Predict the outcome and then try out the following program

// 首先预测一下下面程序的结果，然后动手实验下

package main

import "fmt"

type Base struct {}

func (Base) Magic() {
    fmt.Println("base magic")
}

func (self Base) MoreMagic() {
    self.Magic()
    self.Magic()
}

type Voodoo struct {
    Base
}

func (Voodoo) Magic() {
    fmt.Println("voodoo magic")
}

func main() {
    v := new(Voodoo)
    v.Magic()
    v.MoreMagic()
}

/*
voodoo magic
base magic
base magic
*/
