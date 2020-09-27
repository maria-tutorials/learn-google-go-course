package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int

func main() {
	fmt.Println("# CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("counter", counter)
}
