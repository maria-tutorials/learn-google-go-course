package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go send(c)

	//receive
	receive(c)

	fmt.Println("cee ya")
}

func send(c chan<- int) {
	c <- 1995
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
