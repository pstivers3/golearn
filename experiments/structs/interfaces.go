package main

import (
	"fmt"
	"math"
)

// declare geometry interfaces
type geometry interface {
	area() float64
	perim() float64
}

// declare rectangle struct 
type rectangle struct {
	width, height float64
}
// declare circle struct
type circle struct {
	radius float64
}

// create area method for rectangle struct
func (r rectangle) area() float64 {
	return r.width * r.height
}
// create perimiter  method for rectangle struct
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

// create area method for circlt stuct
func (c circle) area() float64 {
	return math.Pi * c.radius *c.radius
}
// create perimiter method for circlt stuct
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// implement geomentry interfaces; area() and perim()
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimiter: ", g.perim())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
