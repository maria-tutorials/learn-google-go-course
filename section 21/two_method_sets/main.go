package main

import "fmt"

type person struct {
	name string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return fmt.Sprintf("My name is %v", p.name)
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		name: "Princess Peach",
	}

	fmt.Println(p1.speak())
	saySomething(&p1)
}
