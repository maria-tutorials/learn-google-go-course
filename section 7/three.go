package main

import "fmt"

const (
	banana        = 99.9
	car    string = "car"
)

func main() {
	fmt.Printf("%T\t%v\n", banana, banana)
	fmt.Printf("%T\t%v\n", car, car)
}
