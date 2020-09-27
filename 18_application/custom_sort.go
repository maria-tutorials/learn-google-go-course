package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	p1 := Person{"Captain", 72}
	p2 := Person{"Natasha", 32}
	p3 := Person{"Clint", 35}
	p4 := Person{"Maria", 34}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("before", people)

	sort.Sort(ByAge(people))

	fmt.Println("after", people)
}
