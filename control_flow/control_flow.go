package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	foo()

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("Bananas")
}

func bar() {
	fmt.Println("Buh bye")
}
