package main

import (
	"fmt"

	"./dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fluffy := canine{
		name: "Fluffy",
		age:  dog.Years(5),
	}

	fmt.Println(fluffy)
}
