package main

import "fmt"

func main() {

	cs := make(chan<- string, 2) //send-only channel
	cr := make(<-chan string, 2) //receive-only channel

	//fmt.Println(<-cr)
	//cs <- "Princess Peach"

	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	sendOnly()
	receiveOnly()
}

func sendOnly() {
	c := make(chan<- string, 2) //send-only channel

	c <- "Princess"
	c <- "Peach"

	fmt.Printf("%T\n", c)
}

func receiveOnly() {
	c := make(<-chan string, 2) //receive-only channel

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Printf("%T\n", c)
}
