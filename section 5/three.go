package main

import "fmt"

var x int = 42
var y string = "Maria"
var z bool = true

func main() {
	s := fmt.Sprintf("%d\n%v\n%t", x, y, z)
	fmt.Println(s)
}
