package main

import "fmt"

type Record struct {
	Name string
	Age int
}

func main() {
	// Create a pointer to a record
	rec := &Record{Name: "John"}
	fmt.Println("memory location of rec:", rec) // ??
	fmt.Println("main: ", rec.Name)
	fmt.Println("Age was never set and therefore is: ", rec.Age)
}

func changeName(r *Record) {
	r.Name = "Peter"
	fmt.Println("inside changeName: ", r.Name)
}