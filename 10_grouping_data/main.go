package main

import (
	"fmt"
)

func main() {
	//slices()
	maps()
}

func slices() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// Really cool way to remove data from slices at know positions
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func maps() {
	//composite literal
	m := map[string]int{
		"Amália":  3,
		"Eusébio": 2,
	}
	fmt.Println(m)

	fmt.Println(m["Amália"])

	fmt.Println(m["Lontra"])

	// ,ok idiom !!!
	v, ok := m["Lontra"]
	fmt.Println(v)
	fmt.Println(ok)

	// IMPORTANT
	if v, ok := m["Lontra"]; ok {
		fmt.Println(v)
	}
}
