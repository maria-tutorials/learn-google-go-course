package main

import "testing"

func TestMySum(t *testing.T) {
	res := mySum(2, 3)
	if res != 5 {
		t.Error("Expected", 5, "Got", res)
	}
}
