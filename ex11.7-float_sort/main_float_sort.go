package main

import (
    "fmt"
    "./float64"
)

func main() {
    f1 := float64.NewFloat64Array()

    f1.Fill(10)
    fmt.Printf("Before sorting: %s\n", f1)

    float64.Sort(f1)
    fmt.Printf("After soring: %s\n", f1)

    if float64.IsSorted(f1) {
        fmt.Println("The float64 array is sorted!")
    } else {
        fmt.Println("The float64 array is NOT sorted!")
    }
}

/*
Before sorting: { 40.5 96.2 18.3 3.4 61.2 34.9 23.1 55.1 85.1 67.0  }
After soring: { 3.4 18.3 23.1 34.9 40.5 55.1 61.2 67.0 85.1 96.2  }
The float64 array is sorted!
*/
