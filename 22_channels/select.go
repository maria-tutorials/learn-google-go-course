package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send data
	go send(even, odd, quit)

	//receive data back
	receive(even, odd, quit)

	fmt.Println("cee ya aligator")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)

	q <- 0
	close(q)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("even channel sent", v)
		case v := <-o:
			fmt.Println("odd channel sent", v)
		case v := <-q:
			fmt.Println("quit channel sent", v)
			return
		}
	}
}
