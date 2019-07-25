package main

import "fmt"

func main() {
    capitals := map[string]string {"France": "Paris", "Italy":"Rome", "Japan":"Tokyo"}
    for key, _ := range capitals {
        fmt.Println("Map item: Capital of", key, "is", capitals[key])
    }
}
