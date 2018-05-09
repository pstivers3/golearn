package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

// returns an anonymous function
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
	printcaller("printed from called anonymous function") // how does this work?
	// this doesn't work because the named function doesn't take an argument
	// getPrintMessage("printed from called anonymous function")
}
