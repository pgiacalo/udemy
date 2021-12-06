package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base   float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
	getName() string
}

func print(s shape) {
	fmt.Println(s.getName(), s.getArea())
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}
func (t triangle) getName() string {
	return "triangle"
}

func (s square) getName() string {
	return "square"
}
func (s square) getArea() float64 {
	return s.side * s.side
}

func main() {
	t := triangle{
		height: 8,
		base:   10,
	}
	s := square{
		side: 10,
	}
	print(t)
	print(s)
}
