package main

import "fmt"

type Log struct {
    msg	string
}

type Customer struct {
    Name	string
    log	*Log
}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg
}

func (c *Customer) Log() *Log {
    return c.log
}

func main() {
    c := new(Customer)
    c.Name = "Stephen Zhou"
    c.log = new(Log)
    c.log.msg = "1 - Yes we can!"
    fmt.Println(c)
    fmt.Printf("%v, %v\n", *c, *(c.log))
    fmt.Println()

    // shorter
    c2 := &Customer{"Barak Obama", &Log{"1 - No we cann't!"}}
    fmt.Println(c2)
    fmt.Printf("%v, %v\n", *c2, *(c2.log))
    fmt.Println()

    c2.Log().Add("2 - After me the world will be a better place!")
    fmt.Println(c2)
    fmt.Printf("%v, %v\n", *c2, *(c2.log))

    fmt.Println()
    fmt.Println(c2.Log())

    fmt.Println()
    fmt.Println(c2.Log().String())
}

/*
&{Stephen Zhou 0xc42000e1d0}
{Stephen Zhou 0xc42000e1d0}, {1 - Yes we can!}

&{Barak Obama 0xc42000e200}
{Barak Obama 0xc42000e200}, {1 - No we cann't!}

&{Barak Obama 0xc42000e200}
{Barak Obama 0xc42000e200}, {1 - No we cann't!
2 - After me the world will be a better place!}

1 - No we cann't!
2 - After me the world will be a better place!

1 - No we cann't!
2 - After me the world will be a better place!
*/
