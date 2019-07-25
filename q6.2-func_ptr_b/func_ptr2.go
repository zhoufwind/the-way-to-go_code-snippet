// Question 6.2
// What difference if any is there between the following function calls:
// (A) func DoSomething(a *A) {
//  b = a
//  }
// and:
// (B) func DoSomething(a A) {
//  b = &a
//  }

// In case B) a is passed by value, so a copy of the value is made and b is a pointer to
// that copy.
// So a and b refer to different copies of the (same) value of type A.
// If through b in DoSomething the value is changed, a doesn’t see that change, the
// value of a doesn’t change.

package main

import "fmt"

func main() {
    a := 10
    b := CopyValueByValue(a)
    fmt.Printf("a's position: %p, a's value: %d; b's position: %p, b's value: %d\n", &a, a, b, *b)
    a = 20
    fmt.Printf("Change a's value to: %d\n", a)
    fmt.Printf("a's position: %p, a's value: %d; b's position: %p, b's value: %d\n", &a, a, b, *b)
}

func CopyValueByValue(a int) (b *int) {
    b = &a
    return
}

/*
a's position: 0xc420014068, a's value: 10; b's position: 0xc420014070, b's value: 10
Change a's value to: 20
a's position: 0xc420014068, a's value: 20; b's position: 0xc420014070, b's value: 10
*/
