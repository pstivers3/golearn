// 2.2 Struct Types, Willaim Kennedy vid

package main

import "fmt"

type bill struct {
	flag    bool
	counter int64
	pi      float32
}

type alice struct {
	flag    bool
	counter int64
	pi      float32
}

func main() {
	//named types
	var a alice
	var b bill

	//// won't do implicit conversion on named types
	// b = a
	// fmt.Println(b, a)

	// will do explicit conversion on named types
	b = bill(a)
	fmt.Println(b, a)

	// literal type
	/*
		e1 := struct {
			flag    bool
			counter int16
			pi      float32
		}{
			flag:    true,
			counter: 10,
			pi:      3.1416,
		}
	*/

	//// implicit conversion should work with literal type ??
	// b = e1
	//// error: cannot use e1 (type struct { flag bool; counter int16; pi float32 }) as type bill in assignment
	// fmt.Println(b)

}
