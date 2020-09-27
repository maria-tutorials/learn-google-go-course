package main

import "fmt"

func main() {
	fmt.Println("2+3=", mySum(2, 3))
	fmt.Println("5+4", mySum(5, 4))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
