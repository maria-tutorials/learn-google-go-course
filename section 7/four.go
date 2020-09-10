package main

import "fmt"

func main() {
	num := 26
	fmt.Printf("%d\t%b\t%x\n", num, num, num)

	x := num << 1
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
}
