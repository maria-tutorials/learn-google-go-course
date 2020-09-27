package main

import (
	"log"
	"testing"
)

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{40, 1, 1}, 42},
		test{[]int{10, 2, 1, 1}, 14},
		test{[]int{5, 4, 3, 2, 1}, 15},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		res := mySum(v.data...)
		if res != v.answer {
			t.Error("Expected", v.answer, "Got", res)
		}
		log.Println("OK: Expected", v.answer, "Got", res)
	}

}
