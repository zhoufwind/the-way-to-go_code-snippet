// Exercise 10.1
// Define a struct Address and a struct VCard. The latter contains a
// person’s name, a number of addresses, a birth date, a photo. Try to find the right
// data types. Make your own vcard and print its contents.
// Hint: VCard must contain addresses, will they be included as values or as pointers ?
// The 2nd choice is better, consuming less memory. So an Address struct with a name
// and two pointers to addresses could be printed out with %v as:
// {Kersschot 0x126d2b80 0x126d2be0} 

// 定义结构体 Address 和 VCard，后者包含一个人的名字、地址编号、出生日期和图像，试着选择正确的数据类型。构建一个自己的 vcard 并打印它的内容。
// 提示：
// VCard 必须包含住址，它应该以值类型还是以指针类型放在 VCard 中呢？
// 第二种会好点，因为它占用内存少。包含一个名字和两个指向地址的指针的 Address 结构体可以使用 %v 打印：
// {Kersschot 0x126d2b80 0x126d2be0}

package main

import "fmt"

type Address struct {
    house	string
    street	string
    city	string
    country	string
}

type VCard struct {
    name	string
    addr	*Address
    birth	string
    photo	string
}

func main() {
    vc := new(VCard)
    vc.name = "Stephen"
    vc.addr = new(Address)
    vc.addr.house = "Room.703"
    vc.addr.street = "No.2 Boyun Rd."
    vc.addr.city = "Shanghai"
    vc.addr.country = "China"
    vc.birth = "18 Feb."
    vc.photo = "Photo1"

    fmt.Printf("VCard value: %v, position's value: %v\n", *vc, vc)

    fmt.Printf("Address value: %v, position: %p\n", *vc.addr, vc.addr)
}

/*
VCard value: {Stephen 0xc420050100 18 Feb. Photo1}, position's value: &{Stephen 0xc420050100 18 Feb. Photo1}
Address value: {Room.703 No.2 Boyun Rd. Shanghai China}, position: 0xc420050100
*/
