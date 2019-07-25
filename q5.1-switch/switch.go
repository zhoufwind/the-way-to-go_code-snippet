// Question 5.1
// Give the output of the following code snippet:
// k := 6
// switch k {
// case 4: fmt.Println(“was <= 4”); fallthrough;
// case 5: fmt.Println(“was <= 5”); fallthrough;
// case 6: fmt.Println(“was <= 6”); fallthrough;
// case 7: fmt.Println(“was <= 7”); fallthrough;
// case 8: fmt.Println(“was <= 8”); fallthrough;
// default: fmt.Println(“default case”)
// }

package main

import "fmt"

func main() {
    k := 6
    switch k {
        case 4: fmt.Println("was <= 4"); fallthrough;
        case 5: fmt.Println("was <= 5"); fallthrough;
        case 6: fmt.Println("was <= 6"); fallthrough;
        case 7: fmt.Println("was <= 7"); fallthrough;
        case 8: fmt.Println("was <= 8"); fallthrough;
        default: fmt.Println("default case")
    }
}

/*
was <= 6
was <= 7
was <= 8
default case
*/
