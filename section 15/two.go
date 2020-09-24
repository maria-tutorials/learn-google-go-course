package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4}

	r1 := foo(si...)

	r2 := bar(si)

	fmt.Println(r1)
	fmt.Println(r2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
