// Cobbled together from "Go For Devops," Doak 
//     Chap 1, Constructors
// Work left to do: Create a multi-record struct,
//     not just a single record.

package main

import "fmt"

type Record struct {
	Name string
	Age  int
}

func main() {
	MakeRecord("John Doak", 100)
	MakeRecord("Paul Stivers", 102)
	MakeRecord("", 100)
	MakeRecord("Paul Stivers", 102)
}

func MakeRecord(name string, age int){
	rec, err := NewRecord(name, age)
	fmt.Println("record: ", *rec, "\nerror: ", err)
	if err != nil {
		fmt.Println(err, "------\n")
	}
}

func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be the empty string")
	}
	if age <=0 {
		return nil, fmt.Errorf("age cannot be <= 0")
	}
	return &Record{Name: name, Age: age}, nil
}