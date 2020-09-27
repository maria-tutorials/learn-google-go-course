package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("Negative number: %v", n)
		//errors.New("Negative number 👎 ")
	}
	return 42, nil
}
