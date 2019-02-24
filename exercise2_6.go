package main

import "fmt"

const (
	_ = iota
	a = iota + 2018
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
