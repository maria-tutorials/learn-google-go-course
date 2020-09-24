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
}
