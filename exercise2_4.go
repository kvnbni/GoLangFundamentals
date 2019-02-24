package main

import "fmt"

func main() {
	a := 2
	fmt.Printf("%d\t%b\t%x\n", a, a, a)
	b := 2 << 1
	fmt.Printf("%d\t%b\t%x\n", b, b, b)
}
