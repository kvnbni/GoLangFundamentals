package main

import "fmt"

func main() {
	var x []int
	x = []int{1, 2, 3, 4, 5}
	sumfoo := foo(x...)
	sumbar := bar(x)
	fmt.Println(sumfoo)
	fmt.Println(sumbar)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
