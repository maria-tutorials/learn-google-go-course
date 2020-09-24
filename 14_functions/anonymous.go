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
}
