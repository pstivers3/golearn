// 2.3 Pointer - Part 4 (Stack Growth), William Kennedy vid

package main

// import "fmt"

// run with 10 and then with 1024
const size = 10

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
