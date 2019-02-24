package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
	}

	s := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
	}
	fmt.Println(t)
	fmt.Println(s)

	fmt.Println("Truck features", t.fourWheel, t.doors, t.color)
	fmt.Println("Sedan features", s.luxury, s.doors, s.color)
}
