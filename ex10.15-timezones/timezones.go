// Exercise 10.15
// Make an alias type TZ for int. Define a few constants which define timezones like
// UTC and a map which maps these abbreviations to the full name, like:
//  UTC -> “Universal Greenwich time”
// Define a String() method for type TZ which shows the full name of the
// timezone.

// 为 int 定义别名类型 TZ，定义一些常量表示时区，比如 UTC，定义一个 map，它将时区的
// 缩写映射为它的全称，比如：UTC -> "Universal Greenwich time"。为类型 TZ 定义 String() 方法，它输出时区的全称。

package main

import "fmt"

type TZ int

const (
    HOUR	TZ = 60*60
    UTC		TZ = 0 * HOUR
    PDT		TZ = -8 * HOUR
    CST		TZ = 8 * HOUR
)

var timeZones = map[TZ]string{UTC: "Coordinated Universal Time", PDT: "Pacific Daylight Time", CST: "China Standard Time"}

func (tz TZ) String() string {
    for name, value := range timeZones {
        if tz == name {
            return value
        }
    }
    return ""
}

func main() {
    fmt.Println(CST)
    fmt.Println(0 * HOUR)
    fmt.Println(-8 * HOUR)
}

/*
China Standard Time
Coordinated Universal Time
Pacific Daylight Time
*/
