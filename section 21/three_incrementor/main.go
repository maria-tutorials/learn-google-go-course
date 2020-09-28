package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const goroutines int = 100

func main() {
	inc := 0

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			fmt.Println("\t", inc)
			v := inc
			runtime.Gosched()
			v++
			inc = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value", inc)
}
