package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	
	return map[string]int{"x": 1}
}

func main() {
	s := "Go tell it on the mountian. Again and again and again."
	WordCount(s)
	wc.Test(WordCount)
}
