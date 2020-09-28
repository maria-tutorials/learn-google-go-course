package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var inc int

const goroutines int = 100

func main() {

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			mu.Lock()
			fmt.Println("\t", inc)
			v := inc
			v++
			inc = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final value", inc)
}
