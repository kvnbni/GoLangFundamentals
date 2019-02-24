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

	//Store values of type person in map with key as last name
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k, v.first)
		for _, icecream := range v.icecreams {
			fmt.Println(icecream)
		}
	}
}
