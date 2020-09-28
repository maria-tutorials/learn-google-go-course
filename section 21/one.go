package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("start cpu", runtime.NumCPU())
	fmt.Println("stat goroutines", runtime.NumGoroutine())

	wg.Add(2)

	go func() {
		fmt.Println("It's me Mario")
		wg.Done()
	}()
	go func() {
		fmt.Println("It's me Luigi")
		wg.Done()
	}()

	fmt.Println("ongoing cpu", runtime.NumCPU())
	fmt.Println("ongoing goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end goroutines", runtime.NumGoroutine())
	fmt.Println("buh bye")
}
