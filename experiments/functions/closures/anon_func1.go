package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

// returns an anonymous function
// the anonymous func doesn't return anything. It prints a message
func getPrintMessage() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	//named function
	printMessage("printed from named function")

	// anonymous function
	printcaller := getPrintMessage()
	printcaller("printed from called anonymous function") // how does this work ?? q
}
