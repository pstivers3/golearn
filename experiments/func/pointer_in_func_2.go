package main

import (
	"fmt"
)

func main () {
	v := 1
	incr(&v)               // v is now 2
	fmt.Println(incr(&v))  // 3
}

func incr(p *int) int {
	*p++ // increments the value stored in p (an alias for v) 
	return *p
}



