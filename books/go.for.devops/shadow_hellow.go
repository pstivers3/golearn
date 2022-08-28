package main

import "fmt"

var word = "world"

func main() {
	word := "hello"
	fmt.Print(word, " ")
	printOutter()
}

func printOutter() {
	fmt.Println(word)
}
