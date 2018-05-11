package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t) // current time
	fmt.Println(t.Hour()) // current hour
	fmt.Println(t.Weekday()) // current day of the week
	fmt.Println(time.Saturday) // Saturday
	fmt.Println(time.Sunday) // Sunday
}
