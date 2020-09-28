package main

import "fmt"

func main() {
	simpleChannel()
	bufferChannel()
}

func simpleChannel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferChannel() {
	c := make(chan string, 2) //buffered channel

	c <- "Princess"
	c <- "Peach"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
