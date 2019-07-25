package main

func main() {
    f()
}

func f() {
    sum := 0
    for i := 0; i < 1e6; i++ {
        sum += i
    }
}
