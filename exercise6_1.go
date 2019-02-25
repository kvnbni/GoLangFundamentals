package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

//keyword identifier type

//function with identifier foo and returns the type int
func foo() int {
	return 5
}

//function with identifier bar and returns int and a string
func bar() (int, string) {
	return 2, "Kevin"
}
