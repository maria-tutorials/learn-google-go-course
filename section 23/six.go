package main

import "fmt"

func main() {
	c := make(chan int)
	go send100(c)

	for v := range c {
		fmt.Println(v)
	}
}

func send100(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
