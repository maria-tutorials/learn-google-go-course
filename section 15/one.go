package main

import "fmt"

func main() {
	fmt.Println(foo())

	x, y := bar()

	fmt.Printf("%T %v\n %T %v\n", x, x, y, y)
}

func foo() int {
	return 2020
}

func bar() (int, string) {
	return 1970, "unix"
}
