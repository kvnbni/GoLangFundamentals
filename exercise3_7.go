package main

import "fmt"

func main() {
	var x = "Cookie"
	if x == "James Bond" {
		fmt.Println("Yes he is James Bond")
	} else if x == "Money Bond" {
		fmt.Println("No he is Money Bond")
	} else {
		fmt.Println("Nope he is not a Bond")
	}
}
