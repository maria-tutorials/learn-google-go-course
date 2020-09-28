package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

const goroutines int = 100

var inc int64

func main() {
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			fmt.Println(atomic.LoadInt64(&inc))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value", inc)
}
