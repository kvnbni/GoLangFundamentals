package main

import "fmt"

// function f calculates the factorial of a given int
func f(n int) int {
	if n == 0 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	fmt.Println(f(5))
}
