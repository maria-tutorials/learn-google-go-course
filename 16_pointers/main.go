package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & what is the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a // b became an address
	fmt.Println(b)
	fmt.Println(*b) // pointer to an int // * gives you the value stored at an address
	fmt.Println(*&a)

	*b = 43 // affects the value of a
	fmt.Println(a)
}
