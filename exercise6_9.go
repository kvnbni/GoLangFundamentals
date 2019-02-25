package main

import (
	"fmt"
)

func main() {
	var x []int
	x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(x...))
	fmt.Println(even(sum, x))
}

func sum(numbers ...int) int {

	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

//even is a call back function as we pass another function as an argument
func even(f func(numbers ...int) int, xe []int) int {
	var y []int
	for _, v := range xe {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	// now y has all the even numbers
	return f(y...)

}
