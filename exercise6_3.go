package main

import "fmt"

func main() {
	foo()
}

func foo() {
	defer bar()
	fmt.Println("Printed first")
}

func bar() {
	fmt.Println("Printed second")
}
