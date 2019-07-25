// Exercise 4.4
// Integer division by 0 causes the program to crash, a run-time panic occurs (if it is obvious then the
// compiler can detect it)

package main

func main() {
    a, b := 10, 0
    c := a / b	// panic: runtime error: integer divide by zero
    print(c)
}
