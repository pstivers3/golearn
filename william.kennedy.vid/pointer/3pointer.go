// 2.3 Pointer - Part 4 (Stack Growth), William Kennedy vid

package main

// run with 10 and then with 1024
// note, the memory loction of string s 'HELLO' changes each few recursions of stackCopy when the array size is 1024
// because the stack keeps getting resized to hold copies of the larger array

// const size = 10
const size = 1024

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

	// recursive call to stackCopy from within stackCopy
	stackCopy(s, c, a)
}
