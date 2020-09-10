package main

import "fmt"

const (
	y2017 = 2017 + iota
	y2018 = 2017 + iota
	y2019 = 2017 + iota
	y2020 = 2017 + iota
)

func main() {
	fmt.Println(y2017, y2018, y2019, y2020)
}
