package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	fmt.Println("\ntime.Now().Weekday():", time.Now().Weekday())
	fmt.Println("time.Saturday:", time.Saturday)
	fmt.Println("today + 0:", (today + 0)%7)
	fmt.Println("today + 1:", (today + 1)%7)
	// run on Friday, the next line errored, "index out of range," without '%7'. Not sure why ?? q
	fmt.Println("today + 2:", (today + 2)%7, "\n")

	fmt.Println("When's Saturday?")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
