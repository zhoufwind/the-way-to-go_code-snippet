package main

import "fmt"

type Log struct {
    msg	string
}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg + "ONLY IN LOG"
}

type Customer struct {
    Name	string
    Log
}

//func (c *Customer) String() string {
//    return "Name: " + c.Name + "\nLog: " + fmt.Sprintln(c.Log)
//}

func main() {
    c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
    c.Add("2 - After me the world will be a better place!")
    fmt.Println(c.String())

    fmt.Println(c)
}

/*
1 - Yes we can!
2 - After me the world will be a better place!ONLY IN LOG
1 - Yes we can!
2 - After me the world will be a better place!ONLY IN LOG
*/
