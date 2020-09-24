package main

import (
	"fmt"
)

func main() {
	ii := []int{2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(ii...) //... notation
	fmt.Println("The total is", total)
}

/// func (r receiver) funcName(parameter(s)) (return(s)) { code}
func sum(x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index", i, "we are now adding", v, "to the total which is now", sum)
	}
	return sum
}
