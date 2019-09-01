// 2.3 Pointer - Part 2 (Sharing Data), William Kennedy vid

package main

// import "fmt"

func main() {
	// declare var of type int with value of 10.
	count := 10

	// display the value and address of count
	println("count:\tvalue of [", count, "]\taddr of [", &count, "]")

	// pass the value of count to the memory frame for function incrament_value
	increment_value(count)
	println("count:\tvalue of [", count, "]\taddr of [", &count, "]")

	// pass the address of count to funtion increment_mem
	increment_mem(&count)
	println("count:\tvalue of [", count, "]\taddr of [", &count, "]")
}

func increment_value(inc int) {
	// increment inc in the memory frame for this function
	inc++
	println("inc:\tvalue of [", inc, "]\taddr of [", &inc, "]")
}

func increment_mem(inc *int) {
	// increment the value in the memory location stored in inc
	*inc++
	println("inc:\tvalue of [", inc, "]\taddr of [", &inc, "]")
}
