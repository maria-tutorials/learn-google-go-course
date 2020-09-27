package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	create()
	open()
}

func create() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()

	r := strings.NewReader("Bee and PuppyCat")

	io.Copy(f, r)
}

func open() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(string(bs))
}
