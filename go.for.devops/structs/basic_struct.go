package main

import "fmt"

func main() {
	// create record of type struct
	var record = struct {
		Name string
		Age  int
	}{
		Name: "John Doak",
		Age:  100,
	}
	fmt.Println(record)
	fmt.Printf("%s is %d years old\n", record.Name, record.Age)
}
