package main

import "fmt"

func main() {
    str := "GGGGGGGGGGGGGGGGGGGGGGGGG"
    for i := 0; i < 25; i++ {
        fmt.Printf("%02d - %s\n", len(str[0:i+1]), str[0:i+1])
    }
}

/*
01 - G
02 - GG
03 - GGG
04 - GGGG
05 - GGGGG
06 - GGGGGG
07 - GGGGGGG
08 - GGGGGGGG
09 - GGGGGGGGG
10 - GGGGGGGGGG
11 - GGGGGGGGGGG
12 - GGGGGGGGGGGG
13 - GGGGGGGGGGGGG
14 - GGGGGGGGGGGGGG
15 - GGGGGGGGGGGGGGG
16 - GGGGGGGGGGGGGGGG
17 - GGGGGGGGGGGGGGGGG
18 - GGGGGGGGGGGGGGGGGG
19 - GGGGGGGGGGGGGGGGGGG
20 - GGGGGGGGGGGGGGGGGGGG
21 - GGGGGGGGGGGGGGGGGGGGG
22 - GGGGGGGGGGGGGGGGGGGGGG
23 - GGGGGGGGGGGGGGGGGGGGGGG
24 - GGGGGGGGGGGGGGGGGGGGGGGG
25 - GGGGGGGGGGGGGGGGGGGGGGGGG
*/
