package main

import "fmt"
import "time"

var week time.Duration

func main() {
    t := time.Now()
    fmt.Println(t)	// e.g. 2019-07-09 15:23:43.310608 +0800 CST m=+0.000411789
    fmt.Printf("%02d.%02d.%4d\n", t.Month(), t.Day(), t.Year())	// 07.09.2019

    t = time.Now().UTC()
    fmt.Println(t)	// 2019-07-09 07:26:35.003686 +0000 UTC
    fmt.Println(time.Now())

    // calculating times:
    week = 60 * 60 * 24 * 7 * 1e9	// must be in nanosec
    week_from_now := t.Add(week)
    fmt.Println(week_from_now)	// 2019-07-16 07:31:14.854513 +0000 UTC

    // formatting times
    fmt.Println(t.Format(time.RFC822))	// 09 Jul 19 07:32 UTC
    fmt.Println(t.Format(time.ANSIC))	// Tue Jul  9 07:32:41 2019
    fmt.Println(t.Format("2006-01-02T15:04:05.000"))	// 2019-07-09T07:36:43.284

    s := t.Format("20060102")
    fmt.Println(t, "=>", s)	// 2019-07-09 07:38:34.068033 +0000 UTC => 20190709
}

/*
2019-07-20 14:42:59.480023 +0800 CST m=+0.000267622
07.20.2019
2019-07-20 06:42:59.480143 +0000 UTC
2019-07-20 14:42:59.480147 +0800 CST m=+0.000391863
2019-07-27 06:42:59.480143 +0000 UTC
20 Jul 19 06:42 UTC
Sat Jul 20 06:42:59 2019
2019-07-20T06:42:59.480
2019-07-20 06:42:59.480143 +0000 UTC => 20190720
*/
