package main

import "fmt"

func main() {
	a := []byte{'g', 'o', 'l', 'a', 'n', 'b'}
	b := []byte{'g', 'o', 'l', 'a', 'n'}
	fmt.Println(Compare(a, b))
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	// Strings are equal except for possible tail
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // String are equal
}

/*
1
*/
