package main

import "fmt"

// create custome type Record
type Record struct {
	Name string
	Age  int
}

func main() {
	// Use custom type Record
	david := Record{Name: "David Justice", Age: 28}
	sarah := Record{Name: "Sarah Murphy", Age: 28}
	fmt.Printf("%+v\n", david)
	fmt.Printf("%+v\n\n", sarah)

	fmt.Println(david)
}
