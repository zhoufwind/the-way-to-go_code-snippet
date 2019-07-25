// Exercise 6.10
// Study and comprehend the working of the following program
// A function returning another function can be used as a factory function. This can be useful when
// you have to create a number of similar functions: write 1 factory function instead of writing them
// all individually. The following function illustrates this: it returns functions that add a suffix to a
// filename when this is not yet present:
// func MakeAddSuffix(suffix string) func(string) string {
//  return func(name string) string {
//  if !strings.HasSuffix(name, suffix) {
//  return name + suffix
// }
// return name
//  }
// }
// Now we can make functions like: addBmp := MakeAddSuffix(“.bmp”)
//  addJpeg := MakeAddSuffix(“.jpeg”)
// and call them as: addBmp(“file”) // returns: file.bmp
//  addJpeg(“file”) // returns: file.jpeg

package main

import "fmt"
import "strings"

func main() {
    addBmp := MakeAddSuffix(".bmp")
    addJpeg := MakeAddSuffix(".jepg")

    fmt.Printf("%s\n", addBmp("file"))
    fmt.Printf("%s\n", addJpeg("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}

/*
file.bmp
file.jepg
*/
