package main

import "fmt"

type rectangle struct {
	width, height int
}

// area method with receiver type of *rectanble (pointer to rectangle struct)`
func (r *rectangle) area() int {
	return r.width * r.height
}

// perimiter method with reciever type of rectangle
func (r rectangle) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r1 := rectangle{10, 5}
	fmt.Println("rectangle: ", r1)
	fmt.Println("area: ", r1.area())
	fmt.Println("perimiter: ", r1.perim())
}
