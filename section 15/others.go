package main

import (
	"fmt"
)

func main() {
	func(x string) {
		fmt.Println("Bom dia", x)
	}("alegria")

	f := func() {
		fmt.Println(1970)
	}
	f()

	g := crazy()
	fmt.Println(g())
}

func crazy() func() int {
	return func() int {
		return 2020
	}
}
