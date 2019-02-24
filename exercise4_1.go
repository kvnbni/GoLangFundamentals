package main

import "fmt"

func main() {
	var numbers [5]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}

	// range returns both index and value. Since we don't want the index we use _
	for _, v := range numbers {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", numbers)
}
