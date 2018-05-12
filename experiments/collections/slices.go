package main

import (
	"fmt"
)

func main() {
    nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

	s := []string{"a","b","c"}

	// make a copy of s into c
	c := make([]string, len(s))
	copy(c, s)
	c = append(c, "d","e")

	fmt.Println(c)
	fmt.Println(c[2:4]) // print a slice of c, from low to excluding high. Remember index starts at zero. 
	cs := c[2:4] // store slice of c in cs
	fmt.Println(cs[0]) // outputs "c"
}
