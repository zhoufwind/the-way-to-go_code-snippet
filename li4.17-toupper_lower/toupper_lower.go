package main

import "fmt"
import "strings"

func main() {
    var ori string = "Hey, how are you Zac?"
    var lower string
    var upper string

    fmt.Printf("The origin string is: %s\n", ori)
    lower = strings.ToLower(ori)
    fmt.Printf("The lowercase string is: %s\n", lower)
    upper = strings.ToUpper(ori)
    fmt.Printf("The uppercase string is: %s\n", upper)
}

/*
The origin string is: Hey, how are you Zac?
The lowercase string is: hey, how are you zac?
The uppercase string is: HEY, HOW ARE YOU ZAC?
*/
