package main

import "fmt"

func main() {
	// switch statement with no expression by default evaluates to true
	switch {
	case true:
		fmt.Println("should print")
	case false:
		fmt.Println("should not print")
	default:
		fmt.Println("never print")
	}
}
