package main

import "fmt"

type Person struct {
	Name string
	HightFeet int
}

func (p Person) IsTall() string {
	if p.HightFeet >= 6 {
		return "is tall"
	} else {
		return "is not tall"
	}
}

func main() {
	mike := Person {
		HightFeet: 6,
	}
	sue := Person {
		HightFeet: 5,
	}
	fmt.Println("Mike", mike.IsTall())
	fmt.Println("Sue", sue.IsTall())
}
