// Exercise 8.1
// Make a map to hold together the number of the day in the week (1 -> 7) with its name.
// Print them out and test for the presence of tuesday and hollyday.

// 创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。

package main

import "fmt"

func main() {
    weekDay := make(map[int]string, 7)
    weekDay[1] = "Monday"
    weekDay[2] = "Tuesday"
    weekDay[3] = "Wednesday"
    weekDay[4] = "Thursday"
    weekDay[5] = "Friday"
    weekDay[6] = "Saturday"
    weekDay[7] = "Sunday"

    for _, val := range weekDay {
        fmt.Println(val)
    }

    fmt.Printf("Have Tuesday? %t\n", ckHaveValue(weekDay, "Tuesday"))
    fmt.Printf("Have Hollyday? %t\n", ckHaveValue(weekDay, "Hollyday"))
}

func ckHaveValue(weekDay map[int]string, str string) (res bool) {
    for _, val := range weekDay {
        if val == str {
            res = true
            break
        }
    }
    return
}

/*
Thursday
Friday
Saturday
Sunday
Monday
Tuesday
Wednesday
Have Tuesday? true
Have Hollyday? false
*/
