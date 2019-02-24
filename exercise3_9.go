package main

import "fmt"

func main() {
	var favSport = "soccer"

	switch favSport {
	case "ping pong":
		fmt.Println("Close but not close enough")

	case "badminton":
		fmt.Println("hmm")

	case "soccer":
		fmt.Println("You got it")

	}
}
