package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourwheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s.doors)
}
