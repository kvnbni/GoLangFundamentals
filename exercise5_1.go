package main

import "fmt"

type person struct {
	first     string
	last      string
	icecreams []string
}

func main() {
	p1 := person{
		first:     "Frank",
		last:      "Lampard",
		icecreams: []string{`Vanilla`, `Blueberry`},
	}

	p2 := person{
		first:     "John",
		last:      "Terry",
		icecreams: []string{`Chocolate`, `Butter Pecan`},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.icecreams {
		fmt.Printf("%v\t%v\n", v, i)
	}
	fmt.Println(p2.first, p2.last)
	for i, v := range p2.icecreams {
		fmt.Printf("%v\t%v\n", v, i)
	}
}
