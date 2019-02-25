package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, p.last, " and I'm ", p.age, " years old.")
}

func main() {
	p := person{
		first: "Kevin",
		last:  "Benni",
		age:   26,
	}
	p.speak()
}
