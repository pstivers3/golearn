package main

import "fmt"

func main() {
	var c Celsius = 100
	fmt.Println(c)  // 100°C
	fmt.Println(7)  // 7
}

type Celsius float64
func (c Celsius) String() string { // ???
	return fmt.Sprintf("%g°C", c)
}
