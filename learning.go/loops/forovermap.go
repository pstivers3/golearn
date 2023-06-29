package main

import "fmt"

func main() {
	uniqueNames := map[string]bool{"Fred": false, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
