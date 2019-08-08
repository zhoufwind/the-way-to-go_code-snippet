// Exercise 10.14
// Make an alias type Day for int. Define an array of strings with the daynames.
// Define a String() method for type Day which shows the dayname.
// Make an enum const type with iota for all the days of the week (MO, TU, …)

// 为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 
// Day 定义 String() 方法，它输出星期几的名字。使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。

// iota: Elegant Constants in Golang
// Ref: https://splice.com/blog/iota-elegant-constants-golang/

// iota: Golang 中优雅的常量
// Ref: https://segmentfault.com/a/1190000000656284

package main

import "fmt"

type Day int

const (
    MO Day = iota
    TU
    WE
    TH
    FR
    SA
    SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
    return fmt.Sprintf("It's %s", dayName[d])
}

func main() {
    var th Day = 3
    fmt.Println(th)

    var d = SU
    fmt.Println(d)

    fmt.Println(0, MO, 1, TU)
}

/*
It's Thursday
It's Sunday
0 It's Monday 1 It's Tuesday
*/
