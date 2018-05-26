// from go by example

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// this will run first because syncronous call
	f("direct")
	// this will run concurently (asyncronously) with the next goroutine, so will interleave with it
	go f("goroutine")

	go func (msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")

}
