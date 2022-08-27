// 2.2 Struct Types, Willaim Kennedy vid

package main

import "fmt"

type example struct {
	flag    bool
	counter int64
	pi      float32
}

func main() {
	// declare vars e1 and e2 of type example set to their zero values
	var e1 example
	//literal construction. Not prefered for zero value construction but can be done.
	e2 := example{}

	// declare a variable of type exxample and init using a struct literal
	e3 := example{
		flag:    true,
		counter: 10,
		pi:      3.14166666,
	}

	// declare a variable of an anonymous type and init using a struct literal
	// don't need to name it if only using it in one place
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.1416,
	}

	// display the value of e1, e2, e4
	fmt.Printf("%+v\n", e1)
	fmt.Printf("%+v\n", e2)
	fmt.Printf("%+v\n", e4)

	// display the field values of e3
	fmt.Println("flag", e3.flag)
	fmt.Println("counter", e3.counter)
	fmt.Println("pi", e3.pi)
}
