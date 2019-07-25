package main

func main() {
    var a int
    var b int64

    a = 15
    b = a + a	// compiler error: cannot use a + a (type int) as type int64 in assignment
    b = b + 5	// ok: 5 is a constant
}
