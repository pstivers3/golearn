// https://www.youtube.com/watch?v=Kf5LLi7Ti9A&t=12s

package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	p := Person {
		Name: "Jane",
		Age: 37,
	}
	
	if p.IsAdult() {
		fmt.Println(p.Name, "is an adult")
	}
}
