package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//*person is a pointer to the type person
func changeMe(p *person) {
	//changing my first name
	//(*p).first
	p.first = "Kbenni"
}

func main() {
	p := person{
		first: "Kevin",
		last:  "Benni",
		age:   26,
	}
	fmt.Println("My first name is ", p.first)
	fmt.Println("My last name is ", p.last)
	fmt.Println("My age is ", p.age)

	changeMe(&p)

	//Printing my first name after the changeMe function was called
	fmt.Println("My first name now is ", p.first)
}
