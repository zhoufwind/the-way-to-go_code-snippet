package main

import (
    "fmt"
    "time"
)

type myTime struct {
    time.Time	// anonymous field
}

func main() {
    m := myTime{time.Now()}
    // calling existing String method on anonymous Time field
    fmt.Println("Full time now: ", m.String())

    // calling myTime.first10Chars
    fmt.Println("First 10 chars: ", m.first10Chars())
}

func (t myTime) first10Chars() string {
    return t.Time.String()[:10]
}

/*
Full time now:  2019-07-28 13:21:27.944366 +0800 CST m=+0.000264304
First 10 chars:  2019-07-28
*/
