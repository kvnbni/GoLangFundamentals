package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)

	//New slices
	// postion 0 to 4
	fmt.Println(x[:5])
	// position 5 to 9
	fmt.Println(x[5:])
	// postion 2 to 6
	fmt.Println(x[2:7])
	// position 1 to 5
	fmt.Println(x[1:6])
}
