package main

import "fmt"

var x = 42

var y float64

const (
	a float64 = 42.68
	b bool    = true
)

func main() {
	z := "Banana"
	fmt.Printf("%v is of type  %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
	fmt.Printf("%v is of type %T\n", z, z)

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)

	z = "Not banana"

	foo()
}

func foo() {
	fmt.Printf("%v is of type %T\n", y, y)
}
