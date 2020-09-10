package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Won't print :(")
	case (2 == 3):
		fmt.Println("Won't print either :(")
	case true:
		fmt.Println("Will print :D")
		fallthrough
	case (2 == 2):
		fmt.Println("Will also print")
	case true:
		fmt.Println("Will it print??")
	}

	moreSwitch()
}

func ourLoop() {
	x := 1
	for {
		x++
		if x > 10 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("done.")
}

func moreSwitch() {
	fruit := "Cherry"

	switch fruit {
	case "Cherry", "Peach":
		fmt.Println("Summer fruits")
	case "Pear":
		fmt.Println("is pear")
	default:
		fmt.Println("no bueno")
	}
}
