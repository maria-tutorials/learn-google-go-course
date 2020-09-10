package main

import "fmt"

func main() {
	numberThree()
	numberFour()
}

func numberThree() {
	by := 1995
	for by <= 2020 {
		fmt.Println(by)
		by++
	}
}

func numberFour() {
	by := 1995
	for {
		if by <= 2020 {
			fmt.Println(by)
			by++
		} else {
			break
		}
	}
}
