package main

import (
	"fmt"
)

func main() {
	defer late()
	fmt.Println("Hello world")
}

func late() {
	defer func() {
		fmt.Println("Inner defered")
	}()
	fmt.Println("I ran")
}
