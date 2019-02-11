package main

import "fmt"

func main() {
	//short declaration variable
	n, err := fmt.Println("Hello world!!", 42, true) // declares a variable and assigns a value to the variable at the same time
	fmt.Println(n)
	fmt.Println(err)
}
