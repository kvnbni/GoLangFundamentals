package main

import "fmt"

func main() {
	a := 2
	fmt.Println("%d\t%b\t%x", a, a, a)
	b := 2 << 1
	fmt.Println("%d\t%b\t%x", b, b, b)
}
