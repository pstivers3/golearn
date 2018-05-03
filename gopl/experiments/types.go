package main

import "fmt"

func main() {
    type C float64 // delcare type C
	type F float64 // declare type F
	var tc C = 6
	var tf F = 6
	// var add_temps_two_types = tc + tf // errors as expected for incompatible types 
	var add_temps = F(tc) + tf
    fmt.Println(add_temps) // prevents declaration above from erroring for not used
}
