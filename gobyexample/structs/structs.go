package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	p1 := person{"Bob", 20}
	fmt.Println(p1)
	p2 := &person{"Ann", 30}
	fmt.Println(p2)
	p1.age = 21
	fmt.Println(p1)
	p2.name = "Anna"
	fmt.Println(p2)
	fmt.Println(person{name:"Jane", age:31})
	fmt.Println(person{age:40})
}
