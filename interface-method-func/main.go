package main

import (
	"fmt"
	"math"
)

type quad struct {
	side float64
}

func (q quad) area() float64 {
	result := q.side * q.side
	return result
}

type circ struct {
	rad float64
}

func (c circ) area() float64 {
	result := math.Pi * 2 * c.rad
	return result
}

type info interface {
	area() float64
}

func measure(i info) {
	fmt.Println(i.area())
}

func main() {
	x := quad{side: 15.0}

	y := circ{rad: 2.50}

	measure(x)
	measure(y)
}