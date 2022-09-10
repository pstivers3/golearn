package main

import (
	"fmt"
	"strings"
)

// define the interface type Stinger
type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}

// Define method String() with receiver p of type Person, so Person implements Stringer.
func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.First)
}

// StrList is a slice of strings
type StrList []string

// Define method String() with receiver s of type StrList, so StrList also implements Stringer.
func (s StrList) String() string {
	return strings.Join(s, ",")
}

// PrintStringer prints the value of a Stringer in stdout
func PrintStringer(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	john := Person{First: "John", Last: "Doak"}
	var nameList Stringer = StrList{"David", "Sarah"}
	var stringList Stringer = StrList{"Go", "tell", "it", "on", "the", "maountain."}
	PrintStringer(john)       // Prints: Doak,John
	PrintStringer(nameList)   // Prints: David,Sarah
	PrintStringer(nameList)   // Prints: David,Sarah
	PrintStringer(stringList) // Prints: Go,tell,it,on,the,moutain.
}
