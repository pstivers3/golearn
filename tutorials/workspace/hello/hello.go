package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse(stringutil.ToUpper("Hello")))
}
