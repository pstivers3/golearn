package main

import "fmt"

var i interface{}

func main() {

	type Person struct {
		First string
	}

	// i = 3
	// i = "hello world"

	i = Person{First: "John"}

	if v, ok := i.(string); ok {
		fmt.Println(v)
	}

	switch v := i.(type) {
	case int:

		fmt.Printf("i was %d\n", v)

	case string:

		fmt.Printf("i was %s\n", v)

	case float64:

		fmt.Printf("i was %v\n", v)

	case Person, *Person:

		fmt.Printf("i was %v\n", v)

	default:

		// %T will print i's underlying type out
		fmt.Printf("i was an unsupported type %T\n", v)
	}
}
