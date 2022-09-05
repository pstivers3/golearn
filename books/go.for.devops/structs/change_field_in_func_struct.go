package main

import "fmt"

type Record struct {
	Name string
	Age int
}

func main() {
	rec := Record{Name: "John"}
	changeName(rec)
	fmt.Println("main: ", rec.Name)
}

func changeName(r Record) {
	r.Name = "Peter"
	fmt.Println("inside ChangeName: ", r.Name)
}
