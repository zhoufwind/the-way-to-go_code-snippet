package main

import "fmt"
import "time"

func main() {
    i := 0
    for {
        fmt.Printf("zzz %d\n", i)
        i = i + 1
        time.Sleep(time.Duration(1)*time.Second)
    }
}
