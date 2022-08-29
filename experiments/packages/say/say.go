package say

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}

func printWrold() {
	fmt.Println("World")
}

func PrintHelloWorld() {
	PrintHello()
	printWorld()
}
