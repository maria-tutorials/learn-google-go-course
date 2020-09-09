package main

import "fmt"

type measurement int

var x measurement

func main() {
	fmt.Printf("%v : %T\n", x, x)
	x = 365
	fmt.Println(x)

	y := int(x)
	fmt.Printf("%v : %T\n", y, y)
}
