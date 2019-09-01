package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// create p1
	p1 := person{"Bob", 20}
	fmt.Println(p1)
	fmt.Println(&p1) // ??
	fmt.Println("\n")

	// create p2
	p2 := &person{"Ann", 30}
	fmt.Println(p2)
	// print memory location of p2 ??
	fmt.Println(&p2, "\n")

	// update age in p1
	p1.age = 21
	fmt.Println(p1)

	// update name in p2
	p2.name = "Anna"
	fmt.Println(p2)

	// assign name and age
	fmt.Println(person{name: "Jane", age: 31})

	// name remains assigned a zero vallue
	fmt.Println(person{age: 40})
}
