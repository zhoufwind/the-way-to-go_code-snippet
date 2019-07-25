package main

import "fmt"

func main() {
	items := []int{10, 20, 30, 40, 50}
	for ix, _ := range items {
		items[ix] *= 2
	}

	fmt.Println(items)
}
