package main

import "fmt"

type customErr struct {
	message string
}

func (e customErr) Error() string {
	return fmt.Sprintf("an error ocurred: %v", e.message)
}

func main() {
	e := customErr{
		message: "my very own error message",
	}
	foo(e)
}

func foo(e error) {
	fmt.Println("inside foo: ", e.(customErr).message) //access message using assertion
}
