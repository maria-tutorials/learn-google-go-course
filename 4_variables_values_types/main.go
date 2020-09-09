package main

import "fmt"

var x = 42

var y float64

func main() {
	z := "Banana"
	fmt.Printf("%v is of type  %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
	fmt.Printf("%v is of type %T\n", z, z)

	z = "Not banana"

	foo()
}

func foo() {
	fmt.Printf("%v is of type %T\n", y, y)
}
