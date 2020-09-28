package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	//receive many: special for
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("cee ya")
}
