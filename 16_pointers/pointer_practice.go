package main

import "fmt"

func main() {
	x := 0
	fmt.Println("before", x)
	fmt.Println("before", &x)
	foo(&x)
	fmt.Println("after", x)
	fmt.Println("after", &x)
}

func foo(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	*x = 1970 //changing the underlying value in memory of the received address
	fmt.Println(x)
	fmt.Println(*x)
}
