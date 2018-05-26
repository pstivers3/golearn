package main

import "fmt"

func main() {
	func (s string) {
		fmt.Println(s)
	} ("boo")  // passes "boo" to the func. Only allowed in an anonymous func
}
