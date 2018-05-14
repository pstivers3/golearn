package main

import (
	"fmt"
)

func main() {
	a := [5]int{0,1,2,3,4}
	sat := a[0:5]
	fmt.Println(sat)
	sata := append(sat, 5, 6)
	fmt.Println(sata)
	sat[0] = 22
	fmt.Println(sat)
	fmt.Println(a)
}
