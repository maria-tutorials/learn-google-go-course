package main

import (
	"fmt"
	"log"
)

type myError struct {
	lat  string
	long string
	err  error
}

func (m myError) Error() string {
	return fmt.Sprintf("an error occured in: %v %v %v", m.lat, m.long, m.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		merr := fmt.Errorf("Negative number: %v", n)
		return 0, myError{"50.99N", "99.50W", merr}
		//errors.New("Negative number 👎 ")
	}
	return 42, nil
}
