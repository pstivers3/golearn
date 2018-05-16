// from https://gobyexample.com/interfacesD
// https://play.golang.org/p/313UebA3rD

// Interfaces are named collections of method signatures.

package main

import "fmt"
import "math"

// create interface for geometric shapes
type geometry interface {
    area() float64
    perim() float64
}

// create structs of type rect and circle 
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

// To implement an interface, need to implement all the methods in the interface.
// implement geometry interface on rect.
// implement area method for rect
func (r rect) area() float64 {
    return r.width * r.height
}
// implement perim method for rect
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// implement geometry interface on circle
// implement area method for circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
// implement perim method for circle
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here's a generic 'measure' function taking advantage of this to work on any `geometry`.
func measure(g geometry) {
	fmt.Println("geometry: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimiter: ", g.perim())
}

func main() {
	// create instance of rect struct
    r := rect{width: 3, height: 4}
	// create instance of circle struct
    c := circle{radius: 5}

    // The 'circle' and 'rect' struct types both implement the 'geometry' interface
	// so we can use instances of these structs as arguments to 'measure'.
    fmt.Println("rectangle")
	measure(r)
    fmt.Println("\ncircle")
    measure(c)
}

