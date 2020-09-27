package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("nothing.txt")
	if err != nil {
		log.Println("error happened", err)
		log.Fatalln(err) //Println() os.Exit(1)
		panic(err)
	}
	defer f2.Close()

}
