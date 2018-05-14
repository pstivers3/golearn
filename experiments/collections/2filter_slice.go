package main

import "fmt"

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}

func odd_int (num int) bool {
	if num%2 == 1 {
		return true 
	}
	return false 
}

func main() {
	s := []int{0,1,2,3,4,5}
	fmt.Println(Filter(s, odd_int))
}

