package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

//area is the method associated with the value of type circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//area is the method associated with the value of type square
func (sq square) area() float64 {
	return sq.length * sq.length
}

func info(s shape) {
	fmt.Println("The area is ", s.area())
}

func main() {
	c := circle{
		radius: 7,
	}

	sq := square{
		length: 4,
	}

	info(c)
	info(sq)

}
