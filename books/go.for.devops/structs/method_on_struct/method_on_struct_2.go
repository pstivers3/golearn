// From A Tour of Go, https://go.dev/tour/methods/1

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// v is reciever, Abs is method function name
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
