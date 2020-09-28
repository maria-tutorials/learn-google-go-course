// Package acdc does things
package acdc

// Sum is just like all the other sums
func Sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}
