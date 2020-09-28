package main

import "fmt"

func main() {
	simpleChannel()
}

func simpleChannel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
