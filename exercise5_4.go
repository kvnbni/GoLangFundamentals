package main

import "fmt"

func main() {
	s := struct {
		first string
		age   int
	}{
		first: "Kevin",
		age:   26,
	}

	fmt.Println(s)
	fmt.Println("My first name is ", s.first)
	fmt.Println("My age is ", s.age)
}
