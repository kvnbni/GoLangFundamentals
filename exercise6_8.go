package main

import "fmt"

// foo is a function which returns a fuction which returns a string
func foo() func() string {
	return func() string {
		return "Returns this"
	}
}

func main() {
	f := foo()
	fmt.Println(f())

}
