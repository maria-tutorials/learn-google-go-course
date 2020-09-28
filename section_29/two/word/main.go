// Package word provides count functions
package word

import "strings"

// UseCount returns a map with the number of times each word appears in given string s
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in the provided s string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
