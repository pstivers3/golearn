package main

import "fmt"

func main() {
	var record = struct {
		Name string
		Age  int
	}{
		Name: "John Doak",
		Age:  100,
	}

	fmt.Println(record)

}
