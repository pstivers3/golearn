// from Go for DevOps, Doak and Justice

package main

import "fmt"

type Record struct {
	Name string
	Age  int
}

func main() {
	john := Record{Name: "John Doak", Age: 100}
	fmt.Println(john.String())
}

// r is a receiver on type Record.
// String is the method function name.
// String returns a csv (comma delimited fields) representing our record.
func (r Record) String() string {
	return fmt.Sprintf("%s, %d", r.Name, r.Age)
}
