// Package docs to test out go doc
package docs

import "fmt"

// This is some cool documentation
func foo() {
	fmt.Println("Bananas")
}

// Sum does some maths
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
