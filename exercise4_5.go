package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// remove 45,46 and 47 from the slice
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}
