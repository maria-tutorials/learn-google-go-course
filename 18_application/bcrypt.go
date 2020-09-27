package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `hunter2`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	givenPass := `hunterhunter`

	err = bcrypt.CompareHashAndPassword(bs, []byte(givenPass))
	if err != nil {
		fmt.Println("Wrong password m8")
		os.Exit(0)
	}
	fmt.Println("yey login successful")
}
