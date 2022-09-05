// From Go For Devops, Doak, Chap 1, Constructors
// Except pring err instead of return err

package main

import "fmt"

type Record struct {
	Name string
	Age  int
}

func main() {
	rec, err := NewRecord("John Doak", 100)
	fmt.Println("record: ", *rec, "\nerror: ", err)
	if err != nil {
		fmt.Println(err)
	}
}

func NewRecord(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be the empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be <= 0")
	}
	return &Record{Name: name, Age: age}, nil
}
