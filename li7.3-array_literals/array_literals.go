package main

import "fmt"

func main() {
    var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

    for i := 0; i < len(arrKeyValue); i++ {
        fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
    }
}

/*
Person at 0 is 
Person at 1 is 
Person at 2 is 
Person at 3 is Chris
Person at 4 is Ron
*/
