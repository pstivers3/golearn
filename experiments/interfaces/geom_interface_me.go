// write a geometry interface without looking at example, as much as possible

package main

import (
	"fmt"
	"math"
)

type geom interface {
	area() float64
	perim() float64
}

type rect struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length*r.width
}

func (r rect) perim() float64 {
	return (2*r.length + 2*r.width)
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c circle) perim() float64 {
	return (2*math.Pi*c.radius)
}

func params(g geom) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	r := rect{
		length: 2,
		width: 4,
	}

	c := circle{
		radius: 1,
	}

	fmt.Println("rectangle")
	params(r)
	fmt.Println("\ncircle")
	params(c)
}

