package main

import "fmt"
import "strconv"

func main() {
	orig := "ABC"
	//orig := "10086"

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error!\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS := strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

/*
The size of ints is: 64
orig ABC is not an integer - exiting with error!
*/
