package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}
	for i, v := range xy {
		fmt.Println(i, v)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}
}
