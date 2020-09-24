package main

import "fmt"

func main() {
	func(x string) {
		fmt.Println("YEY", x)
	}("Pastel de nata")

	f := func() {
		fmt.Println("Bola de Berlim")
	}
	f()

	x := bar()
	fmt.Printf("%T\n", x)
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 1970
	}
}
