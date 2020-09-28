package main

import (
	"fmt"

	"./acdc"
)

func main() {
	fmt.Println(acdc.Sum(1, 2, 3, 4, 5, 6))
	fmt.Println(acdc.Sum(2, 3))
}
