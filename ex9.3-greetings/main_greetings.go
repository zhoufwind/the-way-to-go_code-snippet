// Exercise 9.3
// Make a program main_greetings.go which can greet the user with “Good
// Day”, or “Good Night”. The different greetings should be in a separate package
// greetings.
// In the same package, make a function IsAM() which returns a bool to indicate
// whether the current time is AM (before 12h) or PM; also make functions
// IsAfternoon() and IsEvening().
// Use this in main_greetings to adapt your greeting.
// (Hint: use the time package)

// 创建一个程序 main_greetings.go 能够和用户说 "Good Day" 或者 "Good Night"。不同的问候应该放到 greetings 包中。
// 在同一个包中创建一个 ISAM 函数返回一个布尔值用来判断当前时间是 AM 还是 PM，同样创建 IsAfternoon 和IsEvening 函数。
// 使用 main_greetings 作出合适的问候(提示：使用 time 包)。

package main

import (
    "fmt"
    "./greetings"
)

func main() {
    name := "Stephen"
    fmt.Println(greetings.GoodDay(name))
    fmt.Println(greetings.GoodNight(name))

    fmt.Printf("Current time: %s\n", greetings.GetTime())

    if greetings.IsAM() {
        fmt.Println("Good morning", name)
    } else if greetings.IsAfternoon() {
        fmt.Println("Good afternoon", name)
    } else if greetings.IsEvening() {
        fmt.Println("Good evening", name)
    } else {
        fmt.Println("Good night", name)
    }
}

/*
Good Day Stephen
Good Night Stephen
Current time: 24 Jul 19 18:01 CST
Good afternoon Stephen

*/
