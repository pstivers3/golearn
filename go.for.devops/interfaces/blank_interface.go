package main

import "fmt"

func main() {
	// i = 3
	i = "hello world"
	// i = Person{First: "John"}

	// if v, ok := i.(string); ok {
	//	fmt.Println(v)
	// }

	switch v := i.(type) {
	case int:
		fmt.Printf("i was %d\n", i)
	case string:
		fmt.Printf("i was %s\n", i)
	case float:
		fmt.Printf("i was %v\n", i)
	case Person, *Person:
		fmt.Printf("i was %v\n", i)
	default:
		// %T will print i's underlying type out
		fmt.Printf("i was an unsupported type %T\n", i)
	}
}
