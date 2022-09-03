package main

import "fmt"

func main() {
	m := map[string]int{"n1": 1, "n2": 2}

	// boolean of if ["n1"] exists is stored in ok
	if value, ok := m["n1"]; ok {
		fmt.Println("value of \"n1\" is ", value)
	} else {
		fmt.Println("value of \"n1\" has unknown value")
	}

	if value, ok := m["n3"]; ok {
		fmt.Println("value of \"n3\" is ", value)
	} else {
		fmt.Println("value of \"n3\" has unknown value")
	}
}
